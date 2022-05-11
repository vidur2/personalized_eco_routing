package main

import (
	"line_integrals_fuel_efficiency/googleApiInteraction"
	"line_integrals_fuel_efficiency/util"
	"line_integrals_fuel_efficiency/util/types"
	"testing"
)

func TestRoutesPreffered(t *testing.T) {
	util.InitClient()
	_, err := googleApiInteraction.RouteInter(types.LatitudeLongitude{Lat: 51.2, Long: 52.2}, types.LatitudeLongitude{Lat: 51.200000002, Long: 52.20000000002}, []types.LatitudeLongitude{{Lat: 51.200000001, Long: 52.20000000002}})
	if err != nil {
		t.Log(err)
	}
}
