package src

import (
	"github.com/robfig/cron/v3"

	"github.com/maulanar/go_asset_tracking_management/app"
	"github.com/maulanar/go_asset_tracking_management/src/asset"
)

func Scheduler() *schedulerUtil {
	if scheduler == nil {
		scheduler = &schedulerUtil{}
		if app.APP_ENV == "local" || app.IS_MAIN_SERVER {
			scheduler.Configure()
		}
		scheduler.isConfigured = true
	}
	return scheduler
}

var scheduler *schedulerUtil

type schedulerUtil struct {
	isConfigured bool
}

func (s *schedulerUtil) Configure() {
	c := cron.New()

	// add scheduler func here, for example :
	c.AddFunc("CRON_TZ=Asia/Jakarta * * * * *", func() {
		asset.JobUpdateAssetValue()
	})

	c.Start()
}
