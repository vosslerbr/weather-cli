package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"weather-cli/internal/geocode"
	"weather-cli/internal/urlgen"
)

func GetWeather(geocodeData geocode.GeocodeResult) WeatherResult {
	url := urlgen.GenerateWeatherURL(geocodeData.Latitude, geocodeData.Longitude, geocodeData.TimeZone)

	// fmt.Println("Weather URL is:", url)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var weatherResJSON WeatherResult

	// ? Since we already defined err in this scope, we'll shadow (overwrite it) here
	err = json.Unmarshal(body, &weatherResJSON) // basically JSON.parse() lol

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Weather results found for %s:\n", geocodeData.Name)

	return weatherResJSON
}
