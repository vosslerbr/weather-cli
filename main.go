package main

import (
	"flag"

	"weather-cli/internal/display"
	"weather-cli/internal/geocode"
	"weather-cli/internal/weather"
)

// TODO should really only serve as entry point and call other fns based on args
func main() {
	locationPtr := flag.String("location", "10001", "a valid ZIP or postal code")

	flag.Parse()

	geocodeResults := geocode.GetGeocodeResults(locationPtr)

	weatherResults := weather.GetWeather(geocodeResults)

	display.ShowResults(weatherResults)

}
