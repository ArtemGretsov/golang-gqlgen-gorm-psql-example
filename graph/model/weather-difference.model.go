package model

import "gorm.io/gorm"

type WeatherDifference struct {
	gorm.Model
	ID int
	Temperature string
	Pressure string
	WeatherID int
}