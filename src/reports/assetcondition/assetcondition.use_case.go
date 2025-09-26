package assetcondition

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/maulanar/go_asset_tracking_management/app"
	"grest.dev/grest"
)

// UseCase returns a UseCaseHandler for expected use case functional.
func UseCase(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	u := UseCaseHandler{
		Ctx:   &ctx,
		Query: url.Values{},
	}
	if len(query) > 0 {
		u.Query = query[0]
	}
	return u
}

// UseCaseHandler provides a convenient interface for AssetCondition use case, use UseCase to access UseCaseHandler.
type UseCaseHandler struct {
	AssetCondition

	// injectable dependencies
	Ctx   *app.Ctx   `json:"-" db:"-" gorm:"-"`
	Query url.Values `json:"-" db:"-" gorm:"-"`
}

// Async return UseCaseHandler with async process.
func (u UseCaseHandler) Async(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	ctx.IsAsync = true
	return UseCase(ctx, query...)
}

// Get returns the list of AssetCondition data.
func (u UseCaseHandler) Get() (ViewData, error) {
	res := ViewData{}

	// check permission
	err := u.Ctx.ValidatePermission("asset_conditions.list")
	if err != nil {
		return res, err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// find data
	u.Query.Add(grest.QueryDisablePagination, "true")
	dMap, err := app.Query().Find(tx, &Asset{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	jByte, err := json.Marshal(dMap)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	var data []Asset
	if err := app.BindJSON(jByte, &data); err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	res = ViewData{
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		Datas:     data,
	}

	return res, err
}
