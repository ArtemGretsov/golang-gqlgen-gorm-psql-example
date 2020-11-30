package model

import "gorm.io/gorm"

type Day struct {
  gorm.Model
  ID         int
  Date       string
  IsFullInfo bool
  Rate       Rate
  Weather    Weather
  Tag []Tag `gorm:"many2many:day_tags;"`
}