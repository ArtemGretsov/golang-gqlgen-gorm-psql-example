package resolvers

import (
	"context"
	"github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/model"
)

func (d dayResolver) Weather(ctx context.Context, obj *model.Day) (*model.Weather, error) {
	var weather *model.Weather
	d.DB.Find(weather, obj.ID)

	return weather, nil
}

func (d dayResolver) Rate(ctx context.Context, obj *model.Day) (*model.Rate, error) {

	return nil, nil
}