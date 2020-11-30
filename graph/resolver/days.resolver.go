package resolver

import (
  "context"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/model"
)

func (r *queryResolver) Days(ctx context.Context, day *string) ([]*model.Day, error) {
  var days []*model.Day

  if day != nil {
    r.DB.Where(&model.Day{Date: *day}).Find(&days)
    return days, nil
  }

  r.DB.Find(&days)
  return days, nil
}