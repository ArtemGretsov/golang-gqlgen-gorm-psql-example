//go:generate go run github.com/99designs/gqlgen --verbose
package resolvers

import (
	"github.com/ArtemGretsov/golang-exchanges-rates/graph/generated"
	_ "github.com/ArtemGretsov/golang-exchanges-rates/graph/model"
)

type Resolver struct {
}

//func (r *Resolver) Mutation() generated.MutationResolver {
//	return &mutationResolver{r}
//}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }