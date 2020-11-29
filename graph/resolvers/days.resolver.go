package resolvers

import (
  "context"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/model"
)

func (r *queryResolver) Days(ctx context.Context) ([]*model.Day, error) {
  var days []*model.Day
  r.DB.Find(&days)

  return days, nil
}