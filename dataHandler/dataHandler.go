package datahandler

import (
	"context"
	"encoding/json"
	"line_integrals_fuel_efficiency/googleApiInteraction"
	"line_integrals_fuel_efficiency/prisma/db"
	"line_integrals_fuel_efficiency/util"

	"github.com/sajari/regression"
	"github.com/valyala/fasthttp"
)

type DataCoord struct {
	ActualSpeed float32
	SpeedLimt   float32
}

type DataCoordReq struct {
	Email        string                                   `json:"email"`
	ActualSpeeds []float32                                `json:"actual_speeds"`
	Positions    []googleApiInteraction.LatitudeLongitude `json:"positions"`
	Token        string                                   `json:"token"`
}

func DataHandler(ctx *fasthttp.RequestCtx) error {
	var dataCoord DataCoordReq
	tokCtx := context.Background()

	err := json.Unmarshal(ctx.Request.Body(), &dataCoord)

	if err != nil {
		return err
	}

	err = util.VerifyToken(dataCoord.Token, tokCtx)

	if err != nil {
		return err
	}

	limits := googleApiInteraction.GetSpeedLimit(dataCoord.Positions)
	client := db.NewClient()
	err = client.Prisma.Connect()

	if err != nil {
		return err
	}

	var realModel regression.Regression
	user, err := client.User.FindFirst(db.User.Email.Equals(dataCoord.Email)).Exec(tokCtx)

	if err != nil {
		client.Prisma.Disconnect()
		return err
	}

	model := user.DumpsModel

	err = json.Unmarshal([]byte(model), &realModel)

	if err != nil {
		realModel = regression.Regression{}
	}

	for idx, data := range limits {
		client.UserData.CreateOne(db.UserData.ActualSpeed.Set(float64(dataCoord.ActualSpeeds[idx])), db.UserData.SpeedLimit.Set(float64(data)), db.UserData.UserEmail.Set(dataCoord.Email))
		realModel.Train(regression.DataPoint(float64(dataCoord.ActualSpeeds[idx]), []float64{float64(data)}))
	}

	for _, data := range user.DataPoints() {
		realModel.Train(regression.DataPoint(data.ActualSpeed, []float64{data.SpeedLimit}))
	}

	err = realModel.Run()

	if err != nil {
		client.Prisma.Disconnect()
		return err
	}

	final, _ := json.Marshal(realModel)

	row := client.User.UpsertOne(db.User.Email.Equals(dataCoord.Email))

	_, err = row.Update(db.User.DumpsModel.Set(string(final))).Exec(tokCtx)

	client.Prisma.Disconnect()

	return err
}
