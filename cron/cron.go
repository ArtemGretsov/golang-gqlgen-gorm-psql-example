package cronjob

import (
  "github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/cron/services"
  "github.com/robfig/cron/v3"
  "gorm.io/gorm"
  "log"
)

func Start(db *gorm.DB)  {
  c := cron.New()
  c.AddFunc("0 12 * * *", func() {
    err := services.SaveDayStatistic(db)

    if err != nil {
      log.Print(err)
    }
  })
}