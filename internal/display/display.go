package display

import (
	"fmt"
	"math"

	"weather-cli/internal/weather"
)

func ShowResults(weather weather.WeatherResult) {
	for i := range weather.Daily.Dates {
		date := weather.Daily.Dates[i]
		unit := weather.Units.Temps
		// weatherCode := weather.Daily.WeatherCodes[i]
		highTemp := int(math.Round(float64(weather.Daily.MaxTemps[i])))
		lowTemp := int(math.Round(float64(weather.Daily.MinTemps[i])))

		fmt.Printf("On %s, the high will be %v%s, and the low will be %v%s\n", date, highTemp, unit, lowTemp, unit)
	}
}
