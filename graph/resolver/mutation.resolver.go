package resolver

import (
  "context"
  "github.com/99designs/gqlgen/graphql"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/generated"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/model"
  "github.com/vektah/gqlparser/v2/gqlerror"
  "gorm.io/gorm"
)

func (m mutationResolver) CreateTag(ctx context.Context, input generated.DayTag) (*model.Tag, error) {
  existDay := model.Day{}
  if m.DB.First(&existDay, input.DayID); existDay.ID == 0 {
    graphql.AddError(ctx, gqlerror.Errorf("Day not found"))
  }

  dayTag := &model.Tag{}
  findErr := m.DB.
     Model(&model.Day{}).
     Select("tags.text, tags.id").
     Joins("left join day_tags on days.id = day_tags.day_id").
     Joins("left join tags on day_tags.tag_id = tags.id").
     Where("tags.text = ?", input.Text).
     Where("days.id = ?", input.DayID).
     Scan(dayTag).
     Error

  if findErr != nil {
    return nil, gqlerror.Errorf("Search tag error")
  }

  if dayTag.ID != 0 {
    return dayTag, nil
  }

  tag := &model.Tag{Text: input.Text}

  err := m.DB.Transaction(func(tx *gorm.DB) error {
    if  err := tx.Where(*tag).FirstOrCreate(tag).Error; err != nil {
      return err
    }

    day := model.Day{ID: input.DayID, Tag: []model.Tag{*tag}}
    if err := tx.Save(&day).Error; err != nil {
      return err
    }

    return nil
  })

  if err != nil {
    graphql.AddError(ctx, gqlerror.Errorf("Creating error"))
  }

  return tag, nil
}