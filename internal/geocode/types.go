package geocode

// ? a structure is a collection of named fields
type GeocodeResult struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	TimeZone  string  `json:"timezone"`
}

type GeocodeResponse struct {
	Results []GeocodeResult `json:"results"`
}
