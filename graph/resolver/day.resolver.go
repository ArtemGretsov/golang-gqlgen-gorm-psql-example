package resolver

import (
  "context"
  "errors"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/generated"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/model"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/utils/calcutil"
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
  weather := &model.Weather{}
  d.DB.Find(weather, weather)
  return weather, nil
}

func (d dayResolver) Rate(ctx context.Context, obj *model.Day) (*model.Rate, error) {
  rate := &model.Rate{DayID: obj.ID}
  d.DB.Find(rate, rate)

  return rate, nil
}

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

func (r rateResolver) Difference(ctx context.Context, obj *model.Rate) (*generated.RateDifference, error) {
  var rate = &model.Rate{ID: obj.ID}
  r.DB.Find(rate, rate)

  if rate.ID == 0 {
    return nil, errors.New("Not found rate")
  }

  var prevRate = &model.Rate{ID: obj.ID - 1}
  r.DB.Find(prevRate, prevRate)

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
  var weather = &model.Weather{ID: obj.ID}
  w.DB.Find(weather, weather)

  if weather.ID == 0 {
    return nil, errors.New("Not found weather")
  }

  var prevWeather = &model.Weather{ID: obj.ID - 1}
  w.DB.Find(prevWeather, prevWeather)

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
