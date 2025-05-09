package main

import (
	"fmt"
	"os"

	"weather-cli/internal/display"
	"weather-cli/internal/geocode"
	"weather-cli/internal/weather"
)

// TODO should really only serve as entry point and call other fns based on args
func main() {
	// args will ALWAYS exist (first one is program name)
	args := os.Args

	// we want to make sure the user passes exactly 1 arg
	if len(args) != 2 {
		fmt.Println("Error: please input a single ZIP code, or a city name (wrap multi-word city names with quotes).")

		os.Exit(1)
	}

	zipOrCity := args[1]

	fmt.Println(zipOrCity)
	geocodeResults := geocode.GetGeocodeResults(zipOrCity)

	weatherResults := weather.GetWeather(geocodeResults.Latitude, geocodeResults.Longitude, geocodeResults.TimeZone, geocodeResults.Name)

	display.ShowResults(weatherResults)
}
