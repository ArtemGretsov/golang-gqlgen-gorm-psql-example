package resolver

import (
  "context"
  "github.com/99designs/gqlgen/graphql"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/generated"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/model"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/utils/calcutil"
  "github.com/vektah/gqlparser/v2/gqlerror"
)

type dayResolver struct{ *Resolver }
type rateResolver struct{ *Resolver }
type weatherResolver struct{ *Resolver }

func (r *Resolver) Day() generated.DayResolver {
  return &dayResolver{r}
}

func (r *Resolver) Rate() generated.RateResolver {
  return &rateResolver{r}
}

func (r *Resolver) Weather() generated.WeatherResolver {
  return &weatherResolver{r}
}

func (d dayResolver) Weather(ctx context.Context, obj *model.Day) (*model.Weather, error) {
  weather := &model.Weather{DayID: obj.ID}
  d.DB.Find(weather, weather)

  if weather.ID == 0 {
    return nil, gqlerror.Errorf("Not found weather for day id %d", obj.ID)
  }

  return weather, nil
}

func (d dayResolver) Rate(ctx context.Context, obj *model.Day) (*model.Rate, error) {
  rate := &model.Rate{DayID: obj.ID}
  d.DB.Find(rate, rate)

  if rate.ID == 0 {
    return nil, gqlerror.Errorf("Not found rate for day id %d", obj.ID)
  }

  return rate, nil
}

func (d dayResolver) Tags(ctx context.Context, obj *model.Day) ([]*model.Tag, error) {
  var tags []*model.Tag

  err := d.DB.
    Model(&model.Day{}).
    Select("tags.text, tags.id").
    Joins("inner join day_tags on days.id = day_tags.day_id").
    Joins("inner join tags on day_tags.tag_id = tags.id").
    Where("days.id = ?", obj.ID).
    Scan(&tags).
    Error

  if err != nil {
    graphql.AddErrorf(ctx, "Tags getting error for day id %d", obj.ID)
  }

  return tags, nil
}

func (r rateResolver) Difference(ctx context.Context, obj *model.Rate) (*generated.RateDifference, error) {
  var rate = &model.Rate{ID: obj.ID}
  r.DB.Find(rate, rate)

  var prevRate = &model.Rate{}
  r.DB.Find(prevRate, model.Rate{ID: obj.ID - 1})

  if prevRate.ID == 0 {
    return &generated.RateDifference{
      Usd: "+0.0000%",
      Eur: "+0.0000%",
    }, nil
  }

  rateDifference := generated.RateDifference{
    Usd: calcutil.CalculateDifference(rate.Usd, prevRate.Usd),
    Eur: calcutil.CalculateDifference(rate.Eur, prevRate.Eur),
  }

  return &rateDifference, nil
}

func (w weatherResolver) Difference(ctx context.Context, obj *model.Weather) (*generated.WeatherDifference, error) {
  var weather = &model.Weather{ID: obj.ID, DayID: obj.DayID}
  w.DB.Find(weather, weather)

  var prevWeather = &model.Weather{}
  w.DB.Find(prevWeather, model.Weather{ID: obj.ID - 1})

  if prevWeather.ID == 0 {
    return &generated.WeatherDifference{
      Pressure: "+0.0000%",
      Temperature: "+0.0000%",
    }, nil
  }

  weatherDifference := generated.WeatherDifference{
    Pressure: calcutil.CalculateDifference(weather.Pressure, prevWeather.Pressure),
    Temperature: calcutil.CalculateDifference(weather.Temperature, prevWeather.Temperature),
  }

  return &weatherDifference, nil
}
