package role

import (
	"errors"
	"net/http"
	"net/url"
	"time"

	"gorm.io/gorm"

	"github.com/maulanar/go_asset_tracking_management/app"
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

// UseCaseHandler provides a convenient interface for Role use case, use UseCase to access UseCaseHandler.
type UseCaseHandler struct {
	Role

	// injectable dependencies
	Ctx   *app.Ctx   `json:"-" db:"-" gorm:"-"`
	Query url.Values `json:"-" db:"-" gorm:"-"`
}

// Async return UseCaseHandler with async process.
func (u UseCaseHandler) Async(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	ctx.IsAsync = true
	return UseCase(ctx, query...)
}

// GetByID returns the Role data for the specified ID.
func (u UseCaseHandler) GetByID(id string) (Role, error) {
	res := Role{}

	// check permission
	err := u.Ctx.ValidatePermission("roles.detail")
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
		key = "name"
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

// Get returns the list of Role data.
func (u UseCaseHandler) Get() (app.ListModel, error) {
	res := app.ListModel{}

	// check permission
	err := u.Ctx.ValidatePermission("roles.list")
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
		err = app.Query().PaginationInfo(tx, &Role{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	// return data count if $per_page set to 0
	if res.Results.PageContext.PerPage == 0 {
		return res, err
	}

	// find data
	data, err := app.Query().Find(tx, &Role{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	res.SetData(data, u.Query)

	// save to cache and return if exists
	app.Cache().Set(cacheKey, res)
	return res, err
}

// Create creates a new data Role with specified parameters.
func (u *UseCaseHandler) Create(p *ParamCreate) error {

	// check permission
	err := u.Ctx.ValidatePermission("roles.create")
	if err != nil {
		return err
	}

	// validasi name kosong
	if p.Name.String == "" {
		return app.Error().New(http.StatusBadRequest, "Field 'name' is required")
	}

	// validate param
	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	// set default value for undefined field
	err = u.setDefaultValue(Role{})
	if err != nil {
		return err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	var existing Role
	err = tx.Model(&Role{}).Where("LOWER(name) = LOWER(?)", p.Name.String).First(&existing).Error
	if err == nil {
		// data ditemukan → duplikat
		return app.Error().New(http.StatusBadRequest, "Role name already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// error selain not found → DB error
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// save data to db
	err = tx.Model(&u).Create(&u).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint())

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("POST", "create", p.ID.String, p)
	return nil
}

// UpdateByID updates the Role data for the specified ID with specified parameters.
func (u UseCaseHandler) UpdateByID(id string, p *ParamUpdate) error {

	// check permission
	err := u.Ctx.ValidatePermission("roles.edit")
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
	err = u.setDefaultValue(old)
	if err != nil {
		return err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// update data on the db
	err = tx.Model(&u).Where("id = ?", old.ID).Updates(u).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint(), old.ID.String)

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("PUT", "Update", old.ID.String, old)
	return nil
}

// PartiallyUpdateByID updates the Role data for the specified ID with specified parameters.
func (u UseCaseHandler) PartiallyUpdateByID(id string, p *ParamPartiallyUpdate) error {

	// check permission
	err := u.Ctx.ValidatePermission("roles.edit")
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
	err = u.setDefaultValue(old)
	if err != nil {
		return err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// update data on the db
	err = tx.Model(&u).Where("id = ?", old.ID).Updates(u).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint(), old.ID.String)

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("PATCH", "Partially Update", old.ID.String, old)
	return nil
}

// DeleteByID deletes the Role data for the specified ID.
func (u UseCaseHandler) DeleteByID(id string, p *ParamDelete) error {
	// check permission
	if err := u.Ctx.ValidatePermission("roles.delete"); err != nil {
		return err
	}

	// get existing data
	old, err := u.GetByID(id)
	if err != nil {
		return err
	}

	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// ✅ soft delete via GORM
	if err := tx.Model(&Role{}).Where("id = ?", old.ID).Update("deleted_at", time.Now().UTC()).Error; err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	app.Cache().Invalidate(u.EndPoint(), old.ID.String)
	go u.Ctx.Hook("DELETE", "delete", old.ID.String, old)

	return nil
}

// setDefaultValue set default value of undefined field when create or update Role data.
func (u *UseCaseHandler) setDefaultValue(old Role) error {

	if !old.ID.Valid {
		u.ID = app.NewNullUUID()
	} else {
		u.ID = old.ID
	}

	// Penentuan kode
	// if u.Name.Valid && u.Name.String != "" {
	// 	// Jika kode dikirim dan berbeda dengan data lama, cek ke DB
	// 	if !old.Name.Valid || u.Name.String != old.Name.String {
	// 		err := app.Common().IsFieldValueExists(u.Ctx, u.EndPoint(), "Name", u.TableName(), "name", u.Name.String)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// 	// Jika kode dikirim dan data lama tidak ada, cek ke DB (sudah tercakup di atas)
	// } else {
	// 	// Jika tidak kirim kode dan data lama ada, gunakan data lama
	// 	if old.Name.Valid && old.Name.String != "" {
	// 		u.Name = old.Name
	// 	} else {
	// 		// Jika tidak kirim kode dan data lama tidak ada, generate baru
	// 		newCode, err := app.Common().GenerateCode(u.Ctx, u.TableName(), "name", u.Name.String)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		u.Name.Set(newCode)
	// 	}
	// }

	if u.Ctx.Action.Method == "POST" {
		if !u.IsActive.Valid {
			u.IsActive.Set(true)
		}

	} else {
	}

	return nil
}
