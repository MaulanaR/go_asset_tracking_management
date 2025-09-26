package distributionassetsperdepartment

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanar/go_asset_tracking_management/app"
)

// REST returns a *RESTAPIHandler.
func REST() *RESTAPIHandler {
	return &RESTAPIHandler{}
}

// RESTAPIHandler provides a convenient interface for DistributionAssetsPerDepartment REST API handler.
type RESTAPIHandler struct {
	UseCase UseCaseHandler
}

// injectDeps inject the dependencies of the DistributionAssetsPerDepartment REST API handler.
func (r *RESTAPIHandler) injectDeps(c *fiber.Ctx) error {
	ctx, ok := c.Locals(app.CtxKey).(*app.Ctx)
	if !ok {
		return app.Error().New(http.StatusInternalServerError, "ctx is not found")
	}
	ctx.FiberCtx = c
	r.UseCase = UseCase(*ctx, app.Query().Parse(c.OriginalURL()))
	return nil
}

// Get is the REST API handler for `GET /api/distribution_assets_per_departments`.
func (r *RESTAPIHandler) Get(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	data, err := r.UseCase.Get()
	if err != nil {
		return app.Error().Handler(c, err)
	}

	return c.Render("distribution_asset_per_department", data)
}
