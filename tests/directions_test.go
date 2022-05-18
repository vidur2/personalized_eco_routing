package tests

import (
	"line_integrals_fuel_efficiency/googleApiInteraction"
	"line_integrals_fuel_efficiency/util"
	"testing"
)

func TestDirections(t *testing.T) {
	util.InitClient()
	route, err := googleApiInteraction.DirectionRequest("710+High+Hampton+Run+Alpharetta+GA", "Johns+Creek+High+School+Johns+Creek+GA")

	if err == nil {
		t.Log(route)
	} else {
		t.Log("Directions error " + err.Error())
	}
}
