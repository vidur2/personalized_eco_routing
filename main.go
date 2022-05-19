package main

/*
#cgo LDFLAGS: /Users/vidurmodgil/Desktop/ProgrammingProjects/line_integrals_fuel_efficiency/lib/libregression.a -ldl
#include "/Users/vidurmodgil/Desktop/ProgrammingProjects/line_integrals_fuel_efficiency/lib/regression.h"
*/
import "C"

import (
	createuser "line_integrals_fuel_efficiency/createUserHandler"
	datahandler "line_integrals_fuel_efficiency/dataHandler"
	deleteuserhandler "line_integrals_fuel_efficiency/deleteUserHandler"
	directionshandler "line_integrals_fuel_efficiency/directionsHandler"
	"line_integrals_fuel_efficiency/util"

	"github.com/valyala/fasthttp"
)

func handler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/get_directions":
		directionshandler.HandleDirections(ctx)
	case "/update_model":
		err := datahandler.DataHandler(ctx)
		handleError(ctx, err)
	case "/create_user":
		err := createuser.HandleCreateUser(ctx)
		handleError(ctx, err)
	case "/delete_user":
		err := deleteuserhandler.HandleDeleteUser(ctx)
		handleError(ctx, err)
	}
}

func handleError(ctx *fasthttp.RequestCtx, err error) {
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.Response.AppendBodyString(err.Error())
	} else {
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
		ctx.Response.AppendBodyString("No Errors")
	}
}

func main() {
	util.InitClient()
	C.print(C.CString("Listening on :8080"))
	fasthttp.ListenAndServe(":8080", handler)
}
