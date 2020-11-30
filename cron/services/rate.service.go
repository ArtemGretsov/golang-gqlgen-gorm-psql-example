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
  urlUsd := fmt.Sprintf(rateUrl + "?base=%s&symbols=RUB", "USD")
  rateUsd := RateResponse{}
  errUsd := reqjson.Get(urlUsd, &rateUsd)

  if errUsd != nil {
    return 0, 0, errUsd
  }

  urlEur := fmt.Sprintf(rateUrl + "?base=%s&symbols=RUB", "EUR")
  rateEur := RateResponse{}
  errEur := reqjson.Get(urlEur, &rateEur)

  if errEur != nil {
    return 0, 0, errEur
  }

  return rateUsd.Rates.RUB, rateEur.Rates.RUB, nil
}