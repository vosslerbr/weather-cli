package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"weather-cli/pkg/urlgen"
)

// ? a structure is a collection of named fields
type Result struct {
	Id        int     `json:"id"` //? this tag tells the unmarshal where to find the data for "Id" from our JSON source
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	TimeZone  string  `json:"timezone"`
}

type GeocodeResponse struct {
	Results []Result `json:"results"`
}

func main() {
	fmt.Println("Running weather CLI")

	// args will ALWAYS exist (first one is program name)
	args := os.Args

	// we want to make sure the user passes exactly 1 arg
	if len(args) != 2 {
		fmt.Println("Error: please input a single ZIP code, or a city name (wrap multi-word city names with quotes).")

		os.Exit(1)
	}

	zipOrCity := args[1]

	fmt.Println(zipOrCity)

	url := urlgen.GenerateGeocodeURL(zipOrCity)

	fmt.Printf("Looking up %s...\n", zipOrCity)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("City found!")

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var resJson GeocodeResponse

	// ? Since we already defined err in this scope, we'll shadow (overwrite it) here
	err = json.Unmarshal(body, &resJson) // basically JSON.parse() lol

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if len(resJson.Results) < 1 {
		fmt.Println("Error: no results found for", zipOrCity)
		os.Exit(1)
	}

	fmt.Println("Geocode results:", resJson)

	url = urlgen.GenerateWeatherURL(resJson.Results[0].Latitude, resJson.Results[0].Longitude, resJson.Results[0].TimeZone)

	fmt.Println("Weather URL is:", url)
}
