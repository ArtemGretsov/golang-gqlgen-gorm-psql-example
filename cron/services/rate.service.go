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

func GetRates() (float64, float64, error) {
  rateUsd := RateResponse{}
  rateEur := RateResponse{}
  urlUsd := fmt.Sprintf(rateUrl + "?base=%s&symbols=RUB", "USD")
  urlEur := fmt.Sprintf(rateUrl + "?base=%s&symbols=RUB", "EUR")

  if err := reqjson.Get(urlUsd, &rateUsd); err != nil {
    return 0, 0, err
  }

  if err := reqjson.Get(urlEur, &rateEur); err != nil {
    return 0, 0, err
  }

  return rateUsd.Rates.RUB, rateEur.Rates.RUB, nil
}