package datahandler

import "line_integrals_fuel_efficiency/googleApiInteraction"

type DataCoord struct {
	ActualSpeed float32
	SpeedLimt   float32
}

type DataCoordReq struct {
	Email        string                                   `json:"email"`
	ActualSpeeds []float32                                `json:"actual_speeds"`
	Positions    []googleApiInteraction.LatitudeLongitude `json:"positions"`
	Token        string                                   `json:"token"`
}

type RegressionInputs struct {
	XValues [][]float32 `json:"x_values"`
	YValues []float32   `json:"y_values"`
}
