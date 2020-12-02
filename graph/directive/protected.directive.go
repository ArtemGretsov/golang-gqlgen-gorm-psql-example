package directive

import (
  "context"
  "github.com/99designs/gqlgen/graphql"
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/generated"
)

func SetDirectives(config *generated.Config)  {
  protected(config)
}

func protected(config *generated.Config) {
  config.Directives.Protected = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
    // place to check the token
    return next(ctx)
  }
}