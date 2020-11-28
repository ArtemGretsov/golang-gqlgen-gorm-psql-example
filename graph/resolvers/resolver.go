//go:generate go run github.com/99designs/gqlgen --verbose
package resolvers

import (
	"github.com/ArtemGretsov/golang-exchanges-rates/graph/generated"
	_ "github.com/ArtemGretsov/golang-exchanges-rates/graph/model"
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
