package services

import (
	"github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/model"
	"gorm.io/gorm"
	"time"
)

func SaveDayStatistic(db *gorm.DB)  {
	date := time.Now().Format("2006-01-02")
	temp, pressure := GetWeather()
	usd, eur := GetRates()

	day := model.Day{
		Date: date,
		Weather: model.Weather{
			Temperature: temp,
			Pressure: pressure,
		},
		Rate: model.Rate{
			Usd: usd,
			Eur: eur,
		},
	}

	db.Create(&day)
}