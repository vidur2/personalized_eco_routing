package main

import (
	directionshandler "line_integrals_fuel_efficiency/directionsHandler"
	"line_integrals_fuel_efficiency/util"

	"github.com/valyala/fasthttp"
)

func handler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/get_directions":
		directionshandler.HandleDirections(ctx)
	case "/update_model":
	}
}

func main() {
	util.InitClient()
	fasthttp.ListenAndServe(":8080", handler)
}
