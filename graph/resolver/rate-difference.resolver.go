package resolver

import (
  "context"
  "errors"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/generated"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/model"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/utils/calcutil"
)

func (r *Resolver) Rate() generated.RateResolver {
  return &rateResolver{r}
}

func (r rateResolver) Difference(ctx context.Context, obj *model.Rate) (*generated.RateDifference, error) {
  var rate = &model.Rate{ID: obj.ID}
  r.DB.First(rate)

  if rate.ID == 0 {
    return nil, errors.New("Not found rate")
  }

  var prevRate = &model.Rate{ID: obj.ID - 1}
  r.DB.First(prevRate)

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

type rateResolver struct{ *Resolver }