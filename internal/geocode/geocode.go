package geocode

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"weather-cli/internal/urlgen"
)

func GetGeocodeResults(zipOrCity string) GeocodeResult {
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

	return resJson.Results[0]
}
