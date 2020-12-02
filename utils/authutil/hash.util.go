package authutil

import (
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/model"
  "github.com/dgrijalva/jwt-go"
  "golang.org/x/crypto/bcrypt"
  "os"
)

func HashPassword(password string) (string, error) {
  bPassword := []byte(password)
  hashedPassword, err := bcrypt.GenerateFromPassword(bPassword, bcrypt.DefaultCost)

  if err != nil {
    return "", err
  }

  return string(hashedPassword), nil
}

func CreateJwtToken(user model.User) (string, error) {
  jwtKey := os.Getenv("JWT_KEY")
  token := jwt.New(jwt.SigningMethodHS256)
  claims := make(jwt.MapClaims)
  claims["id"] = user.ID
  claims["login"] = user.Login
  claims["firstName"] = user.FirstName
  claims["lastName"] = user.LastName
  token.Claims = claims

  return token.SignedString([]byte(jwtKey))
}

func ParseJwtToken(jwtToken string) bool {
  jwtKey := os.Getenv("JWT_KEY")
  token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
    return []byte(jwtKey), nil
  })

  if err != nil || !token.Valid {
    return false
  }

  return true
}