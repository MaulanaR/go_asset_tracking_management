package employee

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/maulanar/go_asset_tracking_management/app"
)

// REST returns a *RESTAPIHandler.
func REST() *RESTAPIHandler {
	return &RESTAPIHandler{}
}

// RESTAPIHandler provides a convenient interface for Employee REST API handler.
type RESTAPIHandler struct {
	UseCase UseCaseHandler
}

// injectDeps inject the dependencies of the Employee REST API handler.
func (r *RESTAPIHandler) injectDeps(c *fiber.Ctx) error {
	ctx, ok := c.Locals(app.CtxKey).(*app.Ctx)
	if !ok {
		return app.Error().New(http.StatusInternalServerError, "ctx is not found")
	}
	ctx.FiberCtx = c
	r.UseCase = UseCase(*ctx, app.Query().Parse(c.OriginalURL()))
	return nil
}

// GetByID is the REST API handler for `GET /api/employees/{id}`.
func (r *RESTAPIHandler) GetByID(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	res, err := r.UseCase.GetByID(c.Params("id"))
	if err != nil {
		return app.Error().Handler(c, err)
	}
	resp := app.ListSingleModel{}
	resp.Ctx = r.UseCase.Ctx
	resp.SetData(res, r.UseCase.Query)

	if r.UseCase.IsFlat() {
		return c.JSON(resp)
	}

	return c.JSON(app.NewJSON(resp).ToStructured().Data)
}

// Get is the REST API handler for `GET /api/employees`.
func (r *RESTAPIHandler) Get(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	res, err := r.UseCase.Get()
	if err != nil {
		return app.Error().Handler(c, err)
	}
	res.SetLink(c)
	if r.UseCase.IsFlat() {
		return c.JSON(res)
	}
	return c.JSON(app.NewJSON(res).ToStructured().Data)
}

// Create is the REST API handler for `POST /api/employees`.
func (r *RESTAPIHandler) Create(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	p := ParamCreate{}
	err = app.NewJSON(c.Body()).ToFlat().Unmarshal(&p)
	if err != nil {
		return app.Error().Handler(c, app.Error().New(http.StatusBadRequest, err.Error()))
	}
	p.Ctx = r.UseCase.Ctx
	p.Query = r.UseCase.Query

	err = r.UseCase.Create(&p)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.Query.Get("is_skip_return") == "true" {
		return c.Status(http.StatusCreated).JSON(map[string]any{"message": "Success"})
	}
	res, err := r.UseCase.GetByID(p.ID.String)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	resp := app.ListSingleModel{}
	resp.Ctx = r.UseCase.Ctx
	resp.SetData(res, r.UseCase.Query)
	if r.UseCase.IsFlat() {
		return c.Status(http.StatusCreated).JSON(resp)
	}
	return c.Status(http.StatusCreated).JSON(app.NewJSON(resp).ToStructured().Data)
}

// UpdateByID is the REST API handler for `PUT /api/employees/{id}`.
func (r *RESTAPIHandler) UpdateByID(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	p := ParamUpdate{}
	err = app.NewJSON(c.Body()).ToFlat().Unmarshal(&p)
	if err != nil {
		return app.Error().Handler(c, app.Error().New(http.StatusBadRequest, err.Error()))
	}
	p.Ctx = r.UseCase.Ctx
	p.Query = r.UseCase.Query

	err = r.UseCase.UpdateByID(c.Params("id"), &p)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.Query.Get("is_skip_return") == "true" {
		return c.JSON(map[string]any{"message": "Success"})
	}
	res, err := r.UseCase.GetByID(c.Params("id"))
	if err != nil {
		return app.Error().Handler(c, err)
	}
	resp := app.ListSingleModel{}
	resp.Ctx = r.UseCase.Ctx
	resp.SetData(res, r.UseCase.Query)

	if r.UseCase.IsFlat() {
		return c.JSON(resp)
	}

	return c.JSON(app.NewJSON(resp).ToStructured().Data)
}

// PartiallyUpdateByID is the REST API handler for `PATCH /api/employees/{id}`.
func (r *RESTAPIHandler) PartiallyUpdateByID(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	p := ParamPartiallyUpdate{}
	err = app.NewJSON(c.Body()).ToFlat().Unmarshal(&p)
	if err != nil {
		return app.Error().Handler(c, app.Error().New(http.StatusBadRequest, err.Error()))
	}
	p.Ctx = r.UseCase.Ctx
	p.Query = r.UseCase.Query

	err = r.UseCase.PartiallyUpdateByID(c.Params("id"), &p)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.Query.Get("is_skip_return") == "true" {
		return c.JSON(map[string]any{"message": "Success"})
	}
	res, err := r.UseCase.GetByID(c.Params("id"))
	if err != nil {
		return app.Error().Handler(c, err)
	}
	resp := app.ListSingleModel{}
	resp.Ctx = r.UseCase.Ctx
	resp.SetData(res, r.UseCase.Query)

	if r.UseCase.IsFlat() {
		return c.JSON(resp)
	}

	return c.JSON(app.NewJSON(resp).ToStructured().Data)
}

// DeleteByID is the REST API handler for `DELETE /api/employees/{id}`.
func (r *RESTAPIHandler) DeleteByID(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	p := ParamDelete{}
	err = app.NewJSON(c.Body()).ToFlat().Unmarshal(&p)
	if err != nil {
		return app.Error().Handler(c, app.Error().New(http.StatusBadRequest, err.Error()))
	}
	p.Ctx = r.UseCase.Ctx
	p.Query = r.UseCase.Query

	err = r.UseCase.DeleteByID(c.Params("id"), &p)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	res := map[string]any{
		"code": http.StatusOK,
		"message": r.UseCase.Ctx.Trans("deleted", map[string]string{
			"employees": p.EndPoint(),
			"id":        c.Params("id"),
		}),
	}
	return c.JSON(res)
}
