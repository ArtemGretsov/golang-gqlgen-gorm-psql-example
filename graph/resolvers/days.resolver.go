package resolvers

import (
	"context"
	"fmt"
	"github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/model"
)

func (r *queryResolver) Days(ctx context.Context) ([]*model.Day, error) {
	fmt.Println(ctx)

	return nil, nil
}