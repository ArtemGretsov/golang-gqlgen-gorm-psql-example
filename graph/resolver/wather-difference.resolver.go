package resolver

import (
  "context"
  "errors"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/generated"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/model"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/utils/calcutil"
)

func (r *Resolver) Weather() generated.WeatherResolver {
  return &weatherResolver{r}
}

func (w weatherResolver) Difference(ctx context.Context, obj *model.Weather) (*generated.WeatherDifference, error) {
  var weather = &model.Weather{ID: obj.ID}
  w.DB.First(weather)

  if weather.ID == 0 {
    return nil, errors.New("Not found weather")
  }

  var prevWeather = &model.Weather{ID: obj.ID - 1}
  w.DB.First(prevWeather)

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

type weatherResolver struct{ *Resolver }