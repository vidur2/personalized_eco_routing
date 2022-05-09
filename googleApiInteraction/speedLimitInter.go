package googleApiInteraction

import (
	"encoding/json"
	"fmt"
	"line_integrals_fuel_efficiency/util"

	"github.com/valyala/fasthttp"
)

const BASE_URI_ROADS = "https://roads.googleapis.com/v1/speedLimits?path="

func GetSpeedLimit(positions []LatitudeLongitude) []float32 {
	var limitsRes SpeedLimitsResponses
	finalUri := BASE_URI_ROADS + parseInputParams(positions) + "&key=" + KEY

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	req.SetRequestURI(finalUri)

	util.Client.Do(req, res)

	err := json.Unmarshal(res.Body(), &limitsRes)

	if err != nil {
		fmt.Println(err)
		return []float32{}
	}
	retSlice := make([]float32, len(limitsRes.SpeedLimits))
	for idx, speedObj := range limitsRes.SpeedLimits {
		retSlice[idx] = float32(speedObj.SpeedLimit)
	}

	return retSlice
}

func parseInputParams(positions []LatitudeLongitude) string {
	retString := ""
	for idx, pos := range positions {
		retString += fmt.Sprintf("%f", pos.Lat) + ", " + fmt.Sprintf("%f", pos.Long)
		if idx != len(positions)-1 {
			retString += "|"
		}
	}

	return retString
}

type SpeedLimitsResponses struct {
	SpeedLimits   []speedInfor         `json:"speedLimits"`
	SnappedPoints []snappedPointsInfor `json:"snappedPoints"`
}

type speedInfor struct {
	PlaceId    string `json:"placeId"`
	SpeedLimit int    `json:"speedLimit"`
	Units      string `json:"units"`
}

type snappedPointsInfor struct {
	Location      LatitudeLongitude `json:"location"`
	OriginalIndex int               `json:"originalIndex"`
	PlaceId       string            `json:"placeId"`
}
