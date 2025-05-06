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
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	TimeZone  string  `json:"timezone"`
}

type GeocodeResponse struct {
	Results          []Result `json:"results"`
	GenerationTimeMS float32  `json:"generationtime_ms"`
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

	url := urlgen.Generate(zipOrCity)

	fmt.Println("the url is:", url)

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

	var resJson GeocodeResponse

	// TODO figure out why this can't be err
	err = json.Unmarshal(body, &resJson) // basically JSON.parse() lol

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(resJson)
}
