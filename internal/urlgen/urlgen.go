package urlgen

import (
	"fmt"
	"net/url"
	"strconv"
)

func GenerateGeocodeURL(zipOrCity *string) string {
	baseURL := "https://geocoding-api.open-meteo.com/v1/search"

	params := url.Values{}
	params.Add("name", *zipOrCity)
	params.Add("count", "10")
	params.Add("language", "en")
	params.Add("format", "json")

	encodedParams := params.Encode()

	fullURL := fmt.Sprintf("%s?%s", baseURL, encodedParams)

	return fullURL
}

func GenerateWeatherURL(lat float64, long float64, timezone string) string {
	baseURL := "https://api.open-meteo.com/v1/forecast"

	params := url.Values{}
	params.Add("latitude", strconv.FormatFloat(lat, 'f', 6, 32))
	params.Add("longitude", strconv.FormatFloat(long, 'f', 6, 32))
	params.Add("daily", "weather_code,temperature_2m_max,temperature_2m_min,sunrise,sunset")
	params.Add("models", "gfs_seamless")
	params.Add("timezone", timezone)
	params.Add("temperature_unit", "fahrenheit")

	encodedParams := params.Encode()

	fullURL := fmt.Sprintf("%s?%s", baseURL, encodedParams)

	return fullURL
}
