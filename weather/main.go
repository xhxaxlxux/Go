package main

import (
	"fmt"

	weather "github.com/xhxaxlxux/Go/weather/weatherInfo"
)

func main() {
	name, weather, temp := weather.WeatherInfo()

	fmt.Printf("現在地: %s 天気: %s 気温: %.1f°C\n", name, weather, temp)
}
