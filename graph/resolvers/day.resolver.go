package resolvers

import (
	"context"
	"fmt"
	"github.com/ArtemGretsov/golang-exchanges-rates/graph/model"
)

func (d dayResolver) Weather(ctx context.Context, obj *model.Day) (*model.Weather, error) {
	fmt.Println("eqweq")

	return nil, nil
}

func (d dayResolver) Rate(ctx context.Context, obj *model.Day) (*model.Rate, error) {
	fmt.Println("eqweq")

	return nil, nil
}