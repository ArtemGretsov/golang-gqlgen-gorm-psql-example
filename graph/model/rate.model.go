package model

import "gorm.io/gorm"

type Rate struct {
  gorm.Model
  ID         int
  Usd        float64
  Eur        float64
  DayID      int
  Difference RateDifference
}