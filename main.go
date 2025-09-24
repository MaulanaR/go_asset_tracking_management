package main

import (
	"embed"
	"os"

	"github.com/maulanar/go_asset_tracking_management/app"
	"github.com/maulanar/go_asset_tracking_management/src"
)

//go:embed all:docs
var f embed.FS

func main() {
	app.Config()
	if len(os.Args) == 2 && os.Args[1] == "update" {
		app.IS_GENERATE_OPEN_API_DOC = true
		src.Router()
		app.OpenAPI().Configure().Generate()
		os.Exit(0)
	}

	app.Logger()
	app.Cache()
	app.Validator()
	app.Translator()
	app.FS()
	app.DB()
	defer app.DB().Close()
	app.Server()

	src.Middleware()
	src.Router()
	app.Server().AddOpenAPIDoc("/api/docs", f)

	app.Server().Fiber.Static("/storages", "./storages")

	src.Migrator()
	src.Seeder()
	src.Scheduler()
	err := app.Server().Start()
	if err != nil {
		app.Logger().Fatal().Err(err).Send()
	}
}
