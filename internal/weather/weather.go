package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"weather-cli/internal/urlgen"
)

func GetWeather(latitude float64, longitude float64, timezone string, locationName string) WeatherResult {
	url := urlgen.GenerateWeatherURL(latitude, latitude, timezone)

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

	fmt.Printf("Weather results found for %s!\n", locationName)

	return weatherResJSON
}
