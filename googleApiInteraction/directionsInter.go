package googleApiInteraction

import (
	"encoding/json"
	"fmt"
	"line_integrals_fuel_efficiency/util"

	"github.com/valyala/fasthttp"
)

const BASE_URI_DIRECTIONS = "https://maps.googleapis.com/maps/api/directions/json?origin=%v&destination=%v&key=AIzaSyBNMIScjW-PYeaEMdjst9oXVq2vbvIehSE&alternatives=true"

func DirectionRequest(origin string, destination string) (DirectionsResponse, error) {
	var directionsResponse DirectionsResponse
	uri := fmt.Sprintf(BASE_URI_DIRECTIONS, origin, destination)
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	req.SetRequestURI(uri)
	req.Header.SetMethod(fasthttp.MethodGet)
	fmt.Println(uri)
	err := util.Client.Do(req, res)

	if err != nil {
		return directionsResponse, err
	} else {
		err := json.Unmarshal(res.Body(), &directionsResponse)

		if err != nil {
			return DirectionsResponse{}, err
		} else {
			return directionsResponse, err
		}
	}
}
