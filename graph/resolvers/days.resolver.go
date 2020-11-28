package resolvers

import (
	"context"
	"fmt"
	"github.com/ArtemGretsov/golang-exchanges-rates/graph/model"
)

func (r *queryResolver) Days(ctx context.Context) ([]*model.Day, error) {
	fmt.Println(ctx)
	day := model.Day{"1", "1990-10-10", true, &model.Weather{}, &model.Rate{} }

	return []*model.Day{&day}, nil
}