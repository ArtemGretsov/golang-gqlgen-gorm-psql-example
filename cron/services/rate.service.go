package services

import (
  "fmt"
  reqjson "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/utils/http"
)

const rateUrl string = "https://api.ratesapi.io/api/latest/"

type RateResponse struct {
  Rates struct {
    RUB float64
  }
}

func GetRates() (float64, float64) {
  urlUsd := fmt.Sprintf(rateUrl + "?base=%s&symbols=RUB", "USD")
  rateUsd := RateResponse{}
  reqjson.Get(urlUsd, &rateUsd)

  urlEur := fmt.Sprintf(rateUrl + "?base=%s&symbols=RUB", "EUR")
  rateEur := RateResponse{}
  reqjson.Get(urlEur, &rateEur)

  return rateUsd.Rates.RUB, rateEur.Rates.RUB
}