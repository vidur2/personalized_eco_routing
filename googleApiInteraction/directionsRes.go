package googleApiInteraction

/*
#cgo LDFLAGS: /Users/vidurmodgil/Desktop/ProgrammingProjects/PersonalizedEcoRouting/line_integrals_fuel_efficiency/lib/libregression.a -ldl
#include "/Users/vidurmodgil/Desktop/ProgrammingProjects/PersonalizedEcoRouting/line_integrals_fuel_efficiency/lib/regression.h"
*/
import "C"

import (
	"line_integrals_fuel_efficiency/prisma/db"
	gascalc "line_integrals_fuel_efficiency/util/gas_calcs"
	"math"
	"strconv"

	"github.com/twpayne/go-polyline"
)

// type DataCoordInput struct {
// 	Lat            float64
// 	Long           float64
// 	Velocity       float64
// 	DeltaElevation float64
// }

type DirectionsResponse struct {
	GeocodedWaypoints []GeocodedElements `json:"geocoded_waypoints"`
	Routes            []RouteElements    `json:"routes"`
}

func (d *DirectionsResponse) GetOptimalRoute(user *db.UserModel) RouteElements {
	// var singleRouteInformation []DataCoordInput
	var chosenRoute RouteElements
	lowScore := float32(0)
	var score float32

	for _, route := range d.Routes {
		score = 0.
		for _, leg := range route.Legs {
			score += leg.getGasConsumptionOverPolyline(user)
		}

		if score <= lowScore {
			lowScore = score
			chosenRoute = route
		}
	}

	return chosenRoute
}

type GeocodedElements struct {
	GeocoderStatus string    `json:"geocoder_status"`
	PlaceId        string    `json:"place_id"`
	Types          [2]string `json:"types"`
}
type RouteElements struct {
	Bounds     boundElements `json:"bounds"`
	Copyrights string        `json:"copyrights"`
	Legs       []LegElements `json:"legs"`
}

type LegElements struct {
	Distance         TextValue         `json:"distance"`
	Duration         TextValue         `json:"duration"`
	EndLocation      LatitudeLongitude `json:"end_location"`
	HtmlInstructions string            `json:"html_instructions"`
	Polyline         PolyLine          `json:"polyline"`
	StartLocation    LatitudeLongitude `json:"start_location"`
	TravelMode       string            `json:"travel_mode"`
	Manuever         string            `json:"maneuver"`
}

func (l *LegElements) getSpeedLimit() float64 {
	return l.Distance.Value * 2.237 / l.Duration.Value
}

func (l *LegElements) getGasConsumptionOverPolyline(user *db.UserModel) float32 {
	// startElevation := 0.
	// endElevation := ElevationInformation{}
	baseMpg := user.FuelEfficiency
	unitsConsumed := float32(0.)

	speedLimit := l.getSpeedLimit()
	speedNotFinal, err := C.predict_regression(C.CString(strconv.FormatFloat(speedLimit, 'f', 'g', 32)), C.CString(user.DumpsModel))
	var userSpeed float64

	if err != nil {
		userSpeed = speedLimit
	} else {
		userSpeed, err = strconv.ParseFloat((C.GoString(speedNotFinal)), 32)
	}

	if err != nil {
		userSpeed = speedLimit
	}

	latLongArray := l.Polyline.decodePolyline()
	startLat := latLongArray[0][0]
	startLong := latLongArray[0][1]
	for idx, latLong := range latLongArray {
		if idx >= 1 {
			lat := latLong[0]
			long := latLong[1]
			norm := math.Sqrt(math.Pow(lat-startLat, 2) + math.Pow(long-startLong, 2))
			unitsConsumed += (1 / gascalc.CalculateMpg(float32(baseMpg), float32(userSpeed))) * float32(norm)
			startLat = lat
			startLong = long
		}

		//endElevation, err = GetElevation(lat, long)
		// if err == nil {
		// 	coords = append(coords, DataCoordInput{
		// 		Lat:      lat,
		// 		Long:     long,
		// 		Velocity: speedLimit,
		// 		// DeltaElevation: endElevation.Results[0].Elevation - startElevation,
		// 	})
		// } else {
		// 	coords = append(coords, DataCoordInput{
		// 		Lat:            lat,
		// 		Long:           long,
		// 		Velocity:       speedLimit,
		// 		DeltaElevation: 0,
		// 	})
		// }
	}

	return unitsConsumed
}

type PolyLine struct {
	Points string `json:"points"`
}

func (p *PolyLine) decodePolyline() [][]float64 {
	coords, _, _ := polyline.DecodeCoords([]byte(p.Points))

	return coords
}

type TextValue struct {
	Text  string  `json:"text"`
	Value float64 `json:"value"`
}

type boundElements struct {
	Northeast LatitudeLongitude `json:"northeast"`
	Southeast LatitudeLongitude `json:"southeast"`
	Northwest LatitudeLongitude `json:"northwest"`
	Southwest LatitudeLongitude `json:"southwest"`
}

type LatitudeLongitude struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"lng"`
}
