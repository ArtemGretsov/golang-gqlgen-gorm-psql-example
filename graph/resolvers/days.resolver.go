package resolvers

import (
	"context"
	"github.com/ArtemGretsov/golang-exchanges-rates/graph/model"
)

func (r *queryResolver) Days(ctx context.Context) ([]*model.Day, error) {
	day := model.Day{"1", "1990-10-10", true, &model.Weather{}, &model.Rate{} }

	return []*model.Day{&day}, nil
}

func (r *queryResolver) Friends(ctx context.Context, obj *model.Day) ([]*model.Day, error) {
	// select * from user where friendid = obj.ID
	return nil,  nil
}