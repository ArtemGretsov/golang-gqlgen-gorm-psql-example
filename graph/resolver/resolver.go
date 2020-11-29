//go:generate go run github.com/99designs/gqlgen --verbose
package resolver

import (
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/generated"
  _ "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/model"
  "gorm.io/gorm"
)

type Resolver struct {
  DB *gorm.DB
}

func (r *Resolver) Query() generated.QueryResolver {
  return &queryResolver{r}
}

func (r *Resolver) Day() generated.DayResolver {
  return &dayResolver{r}
}

type queryResolver struct{ *Resolver }
type dayResolver struct{ *Resolver }
