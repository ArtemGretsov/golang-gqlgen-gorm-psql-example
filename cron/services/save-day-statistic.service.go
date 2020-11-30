package services

import (
  "errors"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/model"
  "gorm.io/gorm"
  "time"
)

func SaveDayStatistic(db *gorm.DB) error  {
  date := time.Now().Format("2006-01-02")
  temp, pressure, errWeather := GetWeather()
  usd, eur, errRate  := GetRates()

  existDay := model.Day{ Date: date }
  db.First(&existDay)

  if existDay.ID != 0 {
    return errors.New("Day exist!")
  }

  if errWeather != nil {
    return errWeather
  }

  if errRate != nil {
    return errRate
  }

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

  return nil
}