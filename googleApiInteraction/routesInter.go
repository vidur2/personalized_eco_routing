package googleApiInteraction

import (
	"encoding/json"
	"fmt"
	"line_integrals_fuel_efficiency/util"
	"line_integrals_fuel_efficiency/util/types"

	"github.com/valyala/fasthttp"
)

const ROUTES_URI = "https://routespreferred.googleapis.com/v1:computeCustomRoutes"

func RouteInter(start types.LatitudeLongitude, end types.LatitudeLongitude, intermediate []types.LatitudeLongitude) (string, error) {
	locations := make([]types.Location, len(intermediate))

	for idx, latLng := range intermediate {
		locations[idx] = types.Location{LatLng: latLng}
	}

	body := types.RoutesPrefferedRequest{
		Origin:        types.Location{LatLng: start},
		Destination:   types.Location{LatLng: end},
		Intermediates: locations,
		RouteObjective: types.RouteObjective{
			RateCard: types.RateCard{
				IncludeTools: false,
			},
		},
	}

	final, err := json.Marshal(body)

	if err != nil {
		return "", err
	}

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetRequestURI(ROUTES_URI + "?key=" + KEY)
	req.SetBody(final)

	err = util.Client.Do(req, res)

	if err != nil {
		return "", err
	}

	fmt.Println(string(res.Body()))

	var parsedRoute DirectionsRes

	err = json.Unmarshal(res.Body(), &parsedRoute)

	if err != nil {
		return "", err
	}

	return parsedRoute.Routes[0].Token, err

}

type DirectionsRes struct {
	Routes        []customRoute `json:"routes"`
	FastestRoute  string        `json:"fastestRoute"`
	ShortestRoute string        `json:"shortestRoute"`
	FallBackInfo  string        `json:"fallbackInfo"`
}

type customRoute struct {
	Route string `json:"route"`
	Token string `json:"token"`
}
