package datahandler

/*
#cgo LDFLAGS: /Users/vidurmodgil/Desktop/ProgrammingProjects/line_integrals_fuel_efficiency/lib/libregression.a -ldl
#include "/Users/vidurmodgil/Desktop/ProgrammingProjects/line_integrals_fuel_efficiency/lib/regression.h"
*/
import "C"

import (
	"context"
	"encoding/json"
	"fmt"
	"line_integrals_fuel_efficiency/googleApiInteraction"
	"line_integrals_fuel_efficiency/prisma/db"
	"line_integrals_fuel_efficiency/util"

	"github.com/valyala/fasthttp"
)

func DataHandler(ctx *fasthttp.RequestCtx) error {
	var dataCoord DataCoordReq
	tokCtx := context.Background()

	err := json.Unmarshal(ctx.Request.Body(), &dataCoord)

	if err != nil {
		return err
	}

	valid, err := util.VerifyToken(dataCoord.Token, tokCtx, dataCoord.Email, util.ClientId(util.ProductionMode))

	if err != nil {
		return err
	} else if !valid {
		return fmt.Errorf("verification error: invalid token information")
	}

	limits := googleApiInteraction.GetSpeedLimit(dataCoord.Positions)
	client := db.NewClient()
	err = client.Prisma.Connect()

	if err != nil {
		return err
	}

	user, err := client.User.FindFirst(db.User.Email.Equals(dataCoord.Email)).Exec(tokCtx)

	if err != nil {
		client.Prisma.Disconnect()
		return err
	}

	xValues := make([][]float32, 0)
	yValues := make([]float32, 0)

	for idx, data := range limits {
		client.UserData.CreateOne(db.UserData.ActualSpeed.Set(float64(dataCoord.ActualSpeeds[idx])), db.UserData.SpeedLimit.Set(float64(data)), db.UserData.UserEmail.Set(dataCoord.Email))
		xValues = append(xValues, []float32{data})
		yValues = append(yValues, dataCoord.ActualSpeeds[idx])
	}

	for _, data := range user.DataPoints() {
		xValues = append(xValues, []float32{float32(data.ActualSpeed)})
		yValues = append(yValues, float32(data.SpeedLimit))
	}

	dataInputs := RegressionInputs{
		XValues: xValues,
		YValues: yValues,
	}

	final, err := json.Marshal(dataInputs)

	if err != nil {
		client.Prisma.Disconnect()
		return err
	}

	model := C.train_regression(C.CString(string(final)))

	if err != nil {
		client.Prisma.Disconnect()
		return err
	}

	row := client.User.UpsertOne(db.User.Email.Equals(dataCoord.Email))

	_, err = row.Create(
		db.User.Email.Set("vmod2005@gmail.com"), db.User.DumpsModel.Set(""), db.User.FuelEfficiency.Set(50),
	).Update(db.User.DumpsModel.Set(C.GoString(model))).Exec(tokCtx)

	client.Prisma.Disconnect()

	return err
}
