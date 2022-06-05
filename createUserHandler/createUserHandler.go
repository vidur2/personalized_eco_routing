package createuser

import (
	"context"
	"encoding/json"
	"fmt"
	"line_integrals_fuel_efficiency/prisma/db"
	"line_integrals_fuel_efficiency/util"
	"line_integrals_fuel_efficiency/util/types"

	"github.com/valyala/fasthttp"
)

type CreateUserOptions struct {
	Username       string  `json:"username"`
	Token          string  `json:"token"`
	FuelEfficiency float64 `json:"fuel_efficiency"`
}

func HandleCreateUser(ctx *fasthttp.RequestCtx) error {
	var createUser CreateUserOptions

	err := json.Unmarshal(ctx.Request.Body(), &createUser)

	if err != nil {
		return err
	}

	prismaCtx := context.Background()
	valid, err := util.VerifyToken(createUser.Token, prismaCtx, createUser.Username, types.ClientId(types.ProductionMode))

	if err != nil {
		return err
	} else if !valid {
		return fmt.Errorf("verification error: invalid token information")
	}

	client := db.NewClient()
	err = client.Prisma.Connect()

	if err != nil {
		return err
	}

	if err != nil {
		client.Prisma.Disconnect()
		return err
	}
	_, err = client.User.CreateOne(db.User.Email.Set(createUser.Username), db.User.DumpsModel.Set(""), db.User.FuelEfficiency.Set(createUser.FuelEfficiency)).Exec(prismaCtx)

	client.Prisma.Disconnect()

	return err
}
