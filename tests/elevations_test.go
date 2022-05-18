package tests

import (
	"line_integrals_fuel_efficiency/googleApiInteraction"
	"line_integrals_fuel_efficiency/util"
	"testing"
)

func TestElevations(t *testing.T) {
	util.InitClient()
	res, err := googleApiInteraction.GetElevation(51.2, 52.2)

	if err == nil {
		t.Log(res)
	} else {
		t.Log("Elevation error " + err.Error())
	}
}
