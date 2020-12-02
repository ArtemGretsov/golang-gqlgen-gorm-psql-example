package resolver

import (
  "context"
  "github.com/99designs/gqlgen/graphql"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/generated"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/model"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/utils/authutil"
  "github.com/vektah/gqlparser/v2/gqlerror"
  "golang.org/x/crypto/bcrypt"
  "gorm.io/gorm"
)

type mutationResolver struct { *Resolver }

func (r* Resolver) Mutation() generated.MutationResolver  {
  return &mutationResolver{r}
}

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

func (m mutationResolver) RegistrationUser(ctx context.Context, input generated.RegistrationUser) (*generated.User, error) {
  const RegistrationError = "Registration error"
  var user model.User
  if err := m.DB.Find(&user, model.User{Login: input.Login}).Error; err != nil {
    return nil, gqlerror.Errorf(RegistrationError)
  }

  if user.ID != 0 {
    return nil, gqlerror.Errorf("User with login '%s' already exist", input.Login)
  }

  hashedPassword, err := authutil.HashPassword(input.Password)

  if err != nil {
    return nil, gqlerror.Errorf(RegistrationError)
  }

  newUser := model.User{
    Password: hashedPassword,
    Login: input.Login,
    FirstName: input.FirstName,
    LastName: input.LastName,
  }

  if err := m.DB.Create(&newUser).Error; err != nil {
    return nil, gqlerror.Errorf(RegistrationError)
  }

  return &generated.User{
    ID: newUser.ID,
    Login: input.Login,
    FirstName: input.FirstName,
    LastName: input.LastName,
  }, nil
}

func (m mutationResolver) LoginUser(ctx context.Context, input *generated.LoginUser) (*generated.User, error) {
  const LoginUserError = "Login user error"
  var user model.User

  if err := m.DB.Find(&user, model.User{Login: input.Login}).Error; err != nil {
    return nil, gqlerror.Errorf(LoginUserError)
  }

  if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
    return nil, gqlerror.Errorf("Incorrect login or password")
  }

  token, err := authutil.CreateJwtToken(user)

  if err != nil {
    return nil, gqlerror.Errorf(LoginUserError)
  }

  return &generated.User{
    Login: input.Login,
    FirstName: user.FirstName,
    LastName: user.LastName,
    ID: user.ID,
    Token: &token,
  }, nil
}
