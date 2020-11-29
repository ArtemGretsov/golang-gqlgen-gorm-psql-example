package model

import "gorm.io/gorm"

type RateDifference struct {
	gorm.Model
	ID  int
	Usd string
	Eur string
	RateID int
}