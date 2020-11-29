package services

import (
	"fmt"
	reqjson "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/utils/http"
	"os"
)

const weatherUrl string = "https://api.openweathermap.org/data/2.5/weather?q=taganrog&units=metric"

type WeatherResponse struct {
	Main struct {
		Temp float64
		Pressure float64
	}
}

func GetWeather() (float64, float64) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	weatherApiKeyUrl := fmt.Sprintf(weatherUrl + "&appid=%s", apiKey)
	weatherResponse := WeatherResponse{}
	reqjson.Get(weatherApiKeyUrl, &weatherResponse)

	return weatherResponse.Main.Temp, weatherResponse.Main.Pressure
}