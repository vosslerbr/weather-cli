package display

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"

	"weather-cli/internal/weather"
)

type Condition struct {
	Day struct {
		Description string `json:"description"`
	} `json:"day"`
	Night struct {
		Description string `json:"description"`
	} `json:"night"`
}

type ConditionsMap map[string]Condition

func ShowResults(weather weather.WeatherResult) {
	data, err := os.ReadFile("./conditions.json")

	if err != nil {
		panic(err)
	}

	var conditionsJSON ConditionsMap

	// ? Since we already defined err in this scope, we'll shadow (overwrite it) here
	err = json.Unmarshal(data, &conditionsJSON) // basically JSON.parse() lol

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for i := range weather.Daily.Dates {
		date := weather.Daily.Dates[i]
		unit := weather.Units.Temps
		highTemp := int(math.Round(float64(weather.Daily.MaxTemps[i])))
		lowTemp := int(math.Round(float64(weather.Daily.MinTemps[i])))

		weatherCode := weather.Daily.WeatherCodes[i]
		condition := conditionsJSON[strconv.Itoa(int(weatherCode))].Day.Description

		fmt.Printf("On %s, it will be %s. The high will be %v%s, and the low will be %v%s\n", date, condition, highTemp, unit, lowTemp, unit)
	}
}
