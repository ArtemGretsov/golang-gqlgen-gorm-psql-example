package resolver

import (
  "context"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/generated"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/model"
)

func (d dayResolver) Weather(ctx context.Context, obj *model.Day) (*model.Weather, error) {
  weather := &model.Weather{DayID: obj.ID}
  d.DB.Where(weather).First(weather)

  return weather, nil
}

func (d dayResolver) Rate(ctx context.Context, obj *model.Day) (*model.Rate, error) {
  rate := &model.Rate{DayID: obj.ID}
  d.DB.Where(rate).First(rate)

  return rate, nil
}

func (r *Resolver) Day() generated.DayResolver {
  return &dayResolver{r}
}

type dayResolver struct{ *Resolver }
