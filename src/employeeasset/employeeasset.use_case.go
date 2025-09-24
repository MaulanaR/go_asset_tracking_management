package employeeasset

import (
	"net/http"
	"net/url"
	"time"

	"github.com/maulanar/go_asset_tracking_management/app"
	"github.com/maulanar/go_asset_tracking_management/src/asset"
	"github.com/maulanar/go_asset_tracking_management/src/attachment"
	"github.com/maulanar/go_asset_tracking_management/src/condition"
	"github.com/maulanar/go_asset_tracking_management/src/employee"
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

// UseCaseHandler provides a convenient interface for EmployeeAsset use case, use UseCase to access UseCaseHandler.
type UseCaseHandler struct {
	EmployeeAsset

	// injectable dependencies
	Ctx   *app.Ctx   `json:"-" db:"-" gorm:"-"`
	Query url.Values `json:"-" db:"-" gorm:"-"`
}

// Async return UseCaseHandler with async process.
func (u UseCaseHandler) Async(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	ctx.IsAsync = true
	return UseCase(ctx, query...)
}

// GetByID returns the EmployeeAsset data for the specified ID.
func (u UseCaseHandler) GetByID(id string) (EmployeeAsset, error) {
	res := EmployeeAsset{}

	// check permission
	err := u.Ctx.ValidatePermission("employee_assets.detail")
	if err != nil {
		return res, err
	}

	// get from cache and return if exists
	cacheKey := u.EndPoint() + "." + id
	app.Cache().Get(cacheKey, &res)
	if res.ID.Valid {
		return res, err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// get from db
	key := "id"
	if !app.Validator().IsValid(id, "uuid") {
		key = "code"
	}
	u.Query.Add(key, id)
	err = app.Query().First(tx, &res, u.Query)
	if err != nil {
		return res, u.Ctx.NotFoundError(err, u.EndPoint(), key, id)
	}

	// save to cache and return if exists
	app.Cache().Set(cacheKey, res)
	return res, err
}

// Get returns the list of EmployeeAsset data.
func (u UseCaseHandler) Get() (app.ListModel, error) {
	res := app.ListModel{}

	// check permission
	err := u.Ctx.ValidatePermission("employee_assets.list")
	if err != nil {
		return res, err
	}
	// get from cache and return if exists
	cacheKey := u.EndPoint() + "?" + u.Query.Encode()
	err = app.Cache().Get(cacheKey, &res)
	if err == nil {
		return res, err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// set pagination info
	res.Results.PageContext.Count,
		res.Results.PageContext.Page,
		res.Results.PageContext.PerPage,
		res.Results.PageContext.PageCount,
		err = app.Query().PaginationInfo(tx, &EmployeeAsset{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	// return data count if $per_page set to 0
	if res.Results.PageContext.PerPage == 0 {
		return res, err
	}

	// find data
	data, err := app.Query().Find(tx, &EmployeeAsset{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	res.SetData(data, u.Query)

	// save to cache and return if exists
	app.Cache().Set(cacheKey, res)
	return res, err
}

// Create creates a new data EmployeeAsset with specified parameters.
func (u UseCaseHandler) Create(p *ParamCreate) error {

	// check permission
	err := u.Ctx.ValidatePermission("employee_assets.create")
	if err != nil {
		return err
	}

	// validate param
	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	// set default value for undefined field
	err = p.setDefaultValue(EmployeeAsset{})
	if err != nil {
		return err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// save data to db
	err = tx.Model(&p).Create(&p).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint())

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("POST", "create", p.ID.String, p)
	return nil
}

// UpdateByID updates the EmployeeAsset data for the specified ID with specified parameters.
func (u UseCaseHandler) UpdateByID(id string, p *ParamUpdate) error {

	// check permission
	err := u.Ctx.ValidatePermission("employee_assets.edit")
	if err != nil {
		return err
	}

	// validate param
	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	// get previous data
	old, err := u.GetByID(id)
	if err != nil {
		return err
	}

	// set default value for undefined field
	err = p.setDefaultValue(old)
	if err != nil {
		return err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// update data on the db
	err = tx.Model(&p).Where("id = ?", old.ID).Updates(p).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint(), old.ID.String)

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("PUT", "Update", old.ID.String, old)
	return nil
}

// PartiallyUpdateByID updates the EmployeeAsset data for the specified ID with specified parameters.
func (u UseCaseHandler) PartiallyUpdateByID(id string, p *ParamPartiallyUpdate) error {

	// check permission
	err := u.Ctx.ValidatePermission("employee_assets.edit")
	if err != nil {
		return err
	}

	// validate param
	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	// get previous data
	old, err := u.GetByID(id)
	if err != nil {
		return err
	}

	// set default value for undefined field
	err = p.setDefaultValue(old)
	if err != nil {
		return err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// update data on the db
	err = tx.Model(&p).Where("id = ?", old.ID).Updates(p).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint(), old.ID.String)

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("PATCH", "Partially Update", old.ID.String, old)
	return nil
}

// DeleteByID deletes the EmployeeAsset data for the specified ID.
func (u UseCaseHandler) DeleteByID(id string, p *ParamDelete) error {

	// check permission
	err := u.Ctx.ValidatePermission("employee_assets.delete")
	if err != nil {
		return err
	}

	// validate param
	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	// get previous data
	old, err := u.GetByID(id)
	if err != nil {
		return err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// update data on the db
	err = tx.Model(&p).Where("id = ?", old.ID).Update("deleted_at", time.Now().UTC()).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint(), old.ID.String)

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("DELETE", "DELETE", old.ID.String, old)
	return nil
}

// setDefaultValue set default value of undefined field when create or update EmployeeAsset data.
func (u *UseCaseHandler) setDefaultValue(old EmployeeAsset) error {
	if !old.ID.Valid {
		u.ID = app.NewNullUUID()
	} else {
		u.ID = old.ID
	}

	// validate AssetID
	key := u.AssetID.String
	if !u.AssetID.Valid || u.AssetID.String == "" {
		key = u.AssetCode.String
	}
	if key != "" {
		ass, err := asset.UseCaseHandler{Ctx: u.Ctx, Query: url.Values{}}.GetByID(key)
		if err != nil {
			return err
		}

		// update status to unavailable
		upAsset := asset.ParamUpdate{}
		upAsset.Status.Set("unavailable")
		err = asset.UseCaseHandler{Ctx: u.Ctx, Query: url.Values{}}.UpdateByID(ass.ID.String, &upAsset)
		if err != nil {
			return err
		}
	}

	// validate EmployeeID
	key = u.EmployeeID.String
	if !u.EmployeeID.Valid || u.EmployeeID.String == "" {
		key = u.EmployeeCode.String
	}
	if key != "" {
		_, err := employee.UseCaseHandler{Ctx: u.Ctx, Query: url.Values{}}.GetByID(key)
		if err != nil {
			return err
		}
	}

	// validate ConditionID
	key = u.ConditionID.String
	if !u.ConditionID.Valid || u.ConditionID.String == "" {
		key = u.ConditionCode.String
	}
	if key != "" {
		_, err := condition.UseCaseHandler{Ctx: u.Ctx, Query: url.Values{}}.GetByID(key)
		if err != nil {
			return err
		}
	}

	// validate AttachmentID
	if u.AttachmentID.Valid && u.AttachmentID.String != "" {
		_, err := attachment.UseCaseHandler{Ctx: u.Ctx, Query: url.Values{}}.GetByID(u.AttachmentID.String)
		if err != nil {
			return err
		}
	}

	u.Date.Set(time.Now())

	return nil
}
