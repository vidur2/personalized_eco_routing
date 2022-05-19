package tests

import (
	"context"
	"fmt"
	"line_integrals_fuel_efficiency/prisma/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrisma(t *testing.T) {
	client := db.NewClient()
	client.Prisma.Connect()
	prismaCtx := context.Background()
	createObj := client.User.CreateOne(db.User.Email.Set("vmod2005@gmail.com"), db.User.DumpsModel.Set(""), db.User.FuelEfficiency.Set(50))
	createObj.Exec(prismaCtx)
	updateObj := client.User.UpsertOne(
		db.User.Email.Equals("vmod2005@gmail.com"),
	).Create(
		db.User.Email.Set("vmod2005@gmail.com"), db.User.DumpsModel.Set(""), db.User.FuelEfficiency.Set(0),
	).Update(
		db.User.FuelEfficiency.Set(49.1),
	)
	_, err := updateObj.Exec((prismaCtx))
	if err != nil {
		fmt.Println(err)
	}

	assert.Nil(t, err)

	_, err = client.User.FindUnique(db.User.Email.Equals("vmod2005@gmail.com")).Delete().Exec(prismaCtx)

	assert.Nil(t, err)

	client.Prisma.Disconnect()
}
