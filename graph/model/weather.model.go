package model

import "gorm.io/gorm"

type Weather struct {
	gorm.Model
	ID          int
	Temperature int
	Pressure    int
	DayID       uint
	Difference  WeatherDifference
}