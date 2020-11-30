package model

import "gorm.io/gorm"

type Weather struct {
  gorm.Model
  ID          int
  Temperature float64
  Pressure    float64
  DayID       int
}