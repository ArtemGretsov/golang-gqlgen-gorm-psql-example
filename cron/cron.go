package cronjob

import (
	"github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/cron/services"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func Start(db *gorm.DB)  {
	c := cron.New()
	c.AddFunc("0 12 * * *", func() {
		services.SaveDayStatistic(db)
	})
}