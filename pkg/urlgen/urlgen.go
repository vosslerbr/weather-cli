package urlgen

import "fmt"

func Generate(zip string) string {
	url := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=10&language=en&format=json", zip)

	return url
}
