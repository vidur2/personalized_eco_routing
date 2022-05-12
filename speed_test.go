package main

import (
	"line_integrals_fuel_efficiency/googleApiInteraction"
	"line_integrals_fuel_efficiency/util"
	"testing"
)

func TestSpeed(t *testing.T) {
	util.InitClient()
	res := googleApiInteraction.GetSpeedLimit([]googleApiInteraction.LatitudeLongitude{{Lat: 38.75807927603043, Long: -9.03741754643809}})
	t.Log(res)
}
