package googleApiInteraction

import (
	"encoding/json"
	"fmt"
	"line_integrals_fuel_efficiency/util"

	"github.com/valyala/fasthttp"
)

type ElevationInformation struct {
	Results []ElevationInformationElement `json:"results"`
	Status  string                        `json:"status"`
}

type ElevationInformationElement struct {
	Elevation  float64 `json:"elevation"`
	Location   string  `json:"location"`
	Resolution float64 `json:"resolution"`
}

const BASE_URI_ELEVATIONS = "https://maps.googleapis.com/maps/api/elevation/json?locations=%v, %v&key=AIzaSyBNMIScjW-PYeaEMdjst9oXVq2vbvIehSE"

func GetElevation(lat float64, long float64) (ElevationInformation, error) {
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	final_uri := fmt.Sprintf(BASE_URI_ELEVATIONS, lat, long)

	req.SetRequestURI(final_uri)

	err := util.Client.Do(req, res)

	var elevInfor ElevationInformation

	if err != nil {
		return elevInfor, err
	} else {
		err := json.Unmarshal(res.Body(), &elevInfor)

		if err != nil {
			return ElevationInformation{}, err
		} else {
			return elevInfor, nil
		}
	}
}
