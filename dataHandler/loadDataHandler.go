package datahandler

import (
	"context"
	"encoding/json"
	"fmt"
	"line_integrals_fuel_efficiency/prisma/db"
	"line_integrals_fuel_efficiency/util"
	"line_integrals_fuel_efficiency/util/types"

	"github.com/valyala/fasthttp"
)

func HandleLoadData(ctx *fasthttp.RequestCtx) {

	fmt.Println(string(ctx.Request.Body()))

	err, valid := _loadDataHandler(ctx)

	var res LoadDataRes

	if err == nil {
		res = LoadDataRes{
			Valid: valid,
		}
	} else {
		res = LoadDataRes{
			Valid: valid,
			Error: err.Error(),
		}
	}

	final, _ := json.Marshal(res)

	fmt.Println(string(final))

	ctx.Response.AppendBody(final)

}

func _loadDataHandler(ctx *fasthttp.RequestCtx) (error, bool) {
	var body LoadDataReq

	err := json.Unmarshal(ctx.Request.Body(), &body)

	if err != nil {
		return err, false
	}

	tokCtx := context.Background()

	valid, err := util.VerifyToken(body.IdToken, tokCtx, body.Email, types.ProductionMode)

	if err != nil {
		return err, false
	}

	if valid {

		client := db.NewClient()

		err := client.Prisma.Connect()

		if err != nil {
			return err, false
		}

		_, err = client.User.FindUnique(db.User.Email.Equals(body.Email)).Exec((tokCtx))

		client.Prisma.Disconnect()

		if err != nil {
			return nil, false
		} else {
			return nil, true
		}
	} else {
		return err, false
	}
}
