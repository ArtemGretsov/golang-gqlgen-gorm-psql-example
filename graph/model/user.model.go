package model

import "gorm.io/gorm"

type User struct {
  *gorm.Model
  ID int
  Login string
  Password string
  FirstName string
  LastName string
}