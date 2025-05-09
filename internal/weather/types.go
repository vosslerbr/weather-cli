package weather

type WeatherResult struct {
	Units struct {
		Temps string `json:"temperature_2m_max"`
	} `json:"daily_units"`
	Daily struct {
		Dates        []string  `json:"time"`
		WeatherCodes []int8    `json:"weather_code"`
		MaxTemps     []float32 `json:"temperature_2m_max"`
		MinTemps     []float32 `json:"temperature_2m_min"`
		Sunrises     []string  `json:"sunrise"`
		Sunsets      []string  `json:"sunset"`
	} `json:"daily"`
}
