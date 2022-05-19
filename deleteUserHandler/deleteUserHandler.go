package deleteUserHandler

import (
	"context"
	"encoding/json"
	"fmt"
	"line_integrals_fuel_efficiency/prisma/db"
	"line_integrals_fuel_efficiency/util"
	"line_integrals_fuel_efficiency/util/types"

	"github.com/valyala/fasthttp"
)

func DeleteUserHandler(ctx *fasthttp.RequestCtx) error {
	idCtx := context.Background()

	var deleteOptions DeleteUserOptions

	err := json.Unmarshal(ctx.Request.Body(), &deleteOptions)

	if err != nil {
		return err
	}

	valid, err := util.VerifyToken(deleteOptions.Idtoken, idCtx, deleteOptions.Email, types.ClientId(types.ProductionMode))

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

	_, err = client.User.FindUnique(db.User.Email.Equals(deleteOptions.Email)).Delete().Exec(idCtx)

	return err
}
