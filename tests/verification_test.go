package tests

import (
	"context"
	"fmt"
	"line_integrals_fuel_efficiency/util"
	"line_integrals_fuel_efficiency/util/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

var token string = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjQ4NmYxNjQ4MjAwNWEyY2RhZjI2ZDkyMTQwMThkMDI5Y2E0NmZiNTYiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJhY2NvdW50cy5nb29nbGUuY29tIiwiYXpwIjoiNjE4MTA0NzA4MDU0LTlyOXMxYzRhbGczNmVybGl1Y2hvOXQ1Mm4zMm42ZGdxLmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwiYXVkIjoiNjE4MTA0NzA4MDU0LTlyOXMxYzRhbGczNmVybGl1Y2hvOXQ1Mm4zMm42ZGdxLmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwic3ViIjoiMTAxNjUxMDI4NzA0NDM0NjI1NjE1IiwiZW1haWwiOiJ2bW9kMjAwNUBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiYXRfaGFzaCI6InF5VWhkOXF5WWFuSzA2TWhWZmI5TVEiLCJpYXQiOjE2NTMxNDI3NDcsImV4cCI6MTY1MzE0NjM0NywianRpIjoiYWZiODMwMzRlNzI2MTQ5YTA2OWU5OThiZTIxMGI4Mzc3MGFiNTQ4YyJ9.ajohbs4KT3MmiNwoQmRn9doQqjLkJUjfiEOV7lM1q2nOVANRGyHviYi105skyff5EumQhjwHJK-bb8gLlsTtvi72iGmotRFBAU5a7lApBDlE4O0Bnd3nv8WrzXXuur5MsFyGY4GkYVZA14l9pj07I4hO4E9p63vdiNsOp4xrxJIViI-avXkiU8_LBpVlCbby68jAWYcq0LrUffSO0gXHjKhFJAG7YyyrgNpgmCg02xx13v4VUVMabLH0VKy7GtrldDeYvSU9yH5FWjTRj8hLvqaVgAgi8592HBlwggKSXyS7stWf63_aj9QVvA3zI7tNKC-tjltGXFABqIsAaAEdgQ"

func TestVerification(t *testing.T) {
	util.InitClient()
	valid, err := util.VerifyToken(token, context.Background(), "vmod2005@gmail.com", types.ClientId(types.TestMode))
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, valid, true, "Token is invalid")
}
