package attachment

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

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

// UseCaseHandler provides a convenient interface for Attachment use case, use UseCase to access UseCaseHandler.
type UseCaseHandler struct {
	Attachment

	// injectable dependencies
	Ctx   *app.Ctx   `json:"-" db:"-" gorm:"-"`
	Query url.Values `json:"-" db:"-" gorm:"-"`
}

// Async return UseCaseHandler with async process.
func (u UseCaseHandler) Async(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	ctx.IsAsync = true
	return UseCase(ctx, query...)
}

// GetByID returns the Attachment data for the specified ID.
func (u UseCaseHandler) GetByID(id string) (Attachment, error) {
	res := Attachment{}

	// check permission
	err := u.Ctx.ValidatePermission("attachments.detail")
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

// Get returns the list of Attachment data.
func (u UseCaseHandler) Get() (app.ListModel, error) {
	res := app.ListModel{}

	// check permission
	err := u.Ctx.ValidatePermission("attachments.list")
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
		err = app.Query().PaginationInfo(tx, &Attachment{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	// return data count if $per_page set to 0
	if res.Results.PageContext.PerPage == 0 {
		return res, err
	}

	// find data
	data, err := app.Query().Find(tx, &Attachment{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	res.SetData(data, u.Query)

	// save to cache and return if exists
	app.Cache().Set(cacheKey, res)
	return res, err
}

// Create creates a new data Attachment with specified parameters.
func (u *UseCaseHandler) Create(p *ParamCreate) error {

	// check permission
	err := u.Ctx.ValidatePermission("attachments.create")
	if err != nil {
		return err
	}

	// set default value for undefined field
	err = u.setDefaultValue(Attachment{})
	if err != nil {
		return err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
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

// UpdateByID updates the Attachment data for the specified ID with specified parameters.
func (u UseCaseHandler) UpdateByID(id string, p *ParamUpdate) error {

	// check permission
	err := u.Ctx.ValidatePermission("attachments.edit")
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

// PartiallyUpdateByID updates the Attachment data for the specified ID with specified parameters.
func (u UseCaseHandler) PartiallyUpdateByID(id string, p *ParamPartiallyUpdate) error {

	// check permission
	err := u.Ctx.ValidatePermission("attachments.edit")
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

// DeleteByID deletes the Attachment data for the specified ID.
func (u UseCaseHandler) DeleteByID(id string, p *ParamDelete) error {

	// check permission
	err := u.Ctx.ValidatePermission("attachments.delete")
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

	// delete file
	path, err := os.Getwd()
	os.Remove(fmt.Sprintf(path+"/storages/%s", old.Name.String))

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint(), old.ID.String)

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("DELETE", "DELETE", old.ID.String, old)
	return nil
}

// setDefaultValue set default value of undefined field when create or update Attachment data.
func (u *UseCaseHandler) setDefaultValue(old Attachment) (err error) {
	if !old.ID.Valid {
		u.ID = app.NewNullUUID()
	} else {
		u.ID = old.ID
	}

	// handle file upload
	u.File, err = u.Ctx.FiberCtx.FormFile("file")
	if err == nil {
		dir := "./storages"
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0775)
			if err != nil {
				return app.Error().New(http.StatusInternalServerError, err.Error())
			}
		}

		// rename & save file
		t := time.Now().Format("2006-01-02-15-04-05")
		extension := filepath.Ext(u.File.Filename)
		newName := fmt.Sprintf("%s%s%s", t, "_", filepath.Base(u.File.Filename))
		newName = strings.Replace(newName, " ", "", -1)

		err = u.Ctx.FiberCtx.SaveFile(u.File, fmt.Sprintf("./storages/%s", newName))
		if err != nil {
			return err
		}
		u.Name.Set(newName)
		u.Path.Set("/storages/" + newName)
		u.Extension.Set(extension)
		u.Url.Set(app.APP_URL + u.Path.String)
	}

	return nil
}
