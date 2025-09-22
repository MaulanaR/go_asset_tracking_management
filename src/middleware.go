package src

import (
	"github.com/maulanar/go_asset_tracking_management/app"
	"github.com/maulanar/go_asset_tracking_management/middleware"
)

func Middleware() *middlewareUtil {
	if mdlwr == nil {
		mdlwr = &middlewareUtil{}
		mdlwr.Configure()
		mdlwr.isConfigured = true
	}
	return mdlwr
}

var mdlwr *middlewareUtil

type middlewareUtil struct {
	isConfigured bool
}

func (*middlewareUtil) Configure() {
	app.Server().AddMiddleware(middleware.Ctx().New)
	app.Server().AddMiddleware(middleware.DB().New)
}
