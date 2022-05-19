package types

type Waypoint struct {
	Via      bool     `json:"via"`
	Location Location `json:"location"`
}

type LatitudeLongitude struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"lng"`
}

type Location struct {
	LatLng LatitudeLongitude `json:"latLng"`
}

type RoutesPrefferedRequest struct {
	Origin         Location       `json:"origin"`
	Destination    Location       `json:"destination"`
	Intermediates  []Location     `json:"intermediates"`
	RouteObjective RouteObjective `json:"routeObjective"`
}

type RouteObjective struct {
	RateCard RateCard `json:"rateCard"`
}

type RateCard struct {
	CostPerMinute MonetaryCost `json:"costPerMinute"`
	CostPerKm     MonetaryCost `json:"costPerKm"`
	IncludeTools  bool         `json:"includeTool"`
}

type MonetaryCost struct {
	Value uint8 `json:"value"`
}
