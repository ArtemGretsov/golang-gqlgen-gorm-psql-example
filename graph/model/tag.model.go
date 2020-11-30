package model

import "gorm.io/gorm"

type Tag struct {
  gorm.Model
  ID int
  Text string
}