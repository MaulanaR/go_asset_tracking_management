package src

import (
	"github.com/maulanar/go_asset_tracking_management/app"
	"github.com/maulanar/go_asset_tracking_management/src/asset"
	"github.com/maulanar/go_asset_tracking_management/src/branch"
	"github.com/maulanar/go_asset_tracking_management/src/category"
	"github.com/maulanar/go_asset_tracking_management/src/condition"
	"github.com/maulanar/go_asset_tracking_management/src/department"
	"github.com/maulanar/go_asset_tracking_management/src/employee"
	// import : DONT REMOVE THIS COMMENT
)

func Migrator() *migratorUtil {
	if migrator == nil {
		migrator = &migratorUtil{}
		migrator.Configure()
		if app.APP_ENV == "local" || app.IS_MAIN_SERVER {
			migrator.Run()
		}
		migrator.isConfigured = true
	}
	return migrator
}

var migrator *migratorUtil

type migratorUtil struct {
	isConfigured bool
}

func (*migratorUtil) Configure() {
	app.DB().RegisterTable("main", department.Department{})
	app.DB().RegisterTable("main", condition.Condition{})
	app.DB().RegisterTable("main", category.Category{})
	app.DB().RegisterTable("main", branch.Branch{})
	app.DB().RegisterTable("main", asset.Asset{})
	app.DB().RegisterTable("main", employee.Employee{})
	// RegisterTable : DONT REMOVE THIS COMMENT
}

func (*migratorUtil) Run() {
	tx, err := app.DB().Conn("main")
	if err != nil {
		app.Logger().Fatal().Err(err).Send()
	} else {
		err = app.DB().MigrateTable(tx, "main", app.Setting{})
	}
	if err != nil {
		app.Logger().Fatal().Err(err).Send()
	}
}
