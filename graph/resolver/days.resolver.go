package resolver

import (
  "context"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/model"
  "github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *queryResolver) Days(ctx context.Context, day *string, limit int, offset int) ([]*model.Day, error) {
  var days []*model.Day

  if day != nil {
    err := r.DB.
      Where(&model.Day{Date: *day}).
      Limit(limit).
      Offset(offset).
      Find(&days).
      Error

    if err != nil {
      return nil, gqlerror.Errorf("Day search error")
    }

    return days, nil
  }

  if err := r.DB.Limit(limit).Offset(offset).Find(&days).Error; err != nil {
    return nil, gqlerror.Errorf("Day search error")
  }

  return days, nil
}