package resolver

import (
  "context"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/generated"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/model"
)

func (d dayResolver) Weather(ctx context.Context, obj *model.Day) (*model.Weather, error) {
  weather := &model.Weather{}
  d.DB.Find(weather, weather)
  return weather, nil
}

func (d dayResolver) Rate(ctx context.Context, obj *model.Day) (*model.Rate, error) {
  rate := &model.Rate{DayID: obj.ID}
  d.DB.Find(rate, rate)

  return rate, nil
}

func (r *Resolver) Day() generated.DayResolver {
  return &dayResolver{r}
}

type dayResolver struct{ *Resolver }

func (d dayResolver) Tags(ctx context.Context, obj *model.Day) ([]*model.Tag, error) {
  var tags []*model.Tag

  d.DB.
    Model(&model.Day{}).
    Select("tags.text, tags.id").
    Joins("left join day_tags on days.id = day_tags.day_id").
    Joins("left join tags on day_tags.tag_id = tags.id").
    Where("days.id = ?", obj.ID).
    Scan(&tags)

  return tags, nil
}

