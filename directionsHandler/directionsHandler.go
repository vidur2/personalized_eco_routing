package directionshandler

import (
	"context"
	"encoding/json"
	"fmt"
	"line_integrals_fuel_efficiency/googleApiInteraction"
	db "line_integrals_fuel_efficiency/prisma/db"
	"line_integrals_fuel_efficiency/util"
	"line_integrals_fuel_efficiency/util/types"

	"github.com/valyala/fasthttp"
)

func HandleDirections(ctx *fasthttp.RequestCtx) error {
	client := db.NewClient()
	err := client.Connect()
	prismaCtx := context.Background()

	if err != nil {
		return err
	}

	var infor types.DirectionsReq

	err = json.Unmarshal(ctx.Request.Body(), &infor)

	if err != nil {
		return err
	}

	err = util.VerifyToken(infor.OauthToken, prismaCtx)

	if err != nil {
		return err
	}

	routeInformation, err := googleApiInteraction.DirectionRequest(infor.Start, infor.End)

	if err != nil {
		fmt.Println(err)
		return err
	}
	user, err := client.User.FindFirst(db.User.Email.Equals(infor.User)).Exec(prismaCtx)

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	coords := routeInformation.GetOptimalRoute(user)

	final, err := json.Marshal(coords)

	if err != nil {
		return err
	}
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.AppendBody(final)

	client.Disconnect()

	return nil
}
