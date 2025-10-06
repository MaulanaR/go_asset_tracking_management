package asset

import (
	"net/http"
	"net/url"
	"time"

	"github.com/maulanar/go_asset_tracking_management/app"
	"github.com/maulanar/go_asset_tracking_management/src/attachment"
	"github.com/maulanar/go_asset_tracking_management/src/category"
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

// UseCaseHandler provides a convenient interface for Asset use case, use UseCase to access UseCaseHandler.
type UseCaseHandler struct {
	Asset

	// injectable dependencies
	Ctx   *app.Ctx   `json:"-" db:"-" gorm:"-"`
	Query url.Values `json:"-" db:"-" gorm:"-"`
}

// Async return UseCaseHandler with async process.
func (u UseCaseHandler) Async(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	ctx.IsAsync = true
	return UseCase(ctx, query...)
}

// GetByID returns the Asset data for the specified ID.
func (u UseCaseHandler) GetByID(id string) (Asset, error) {
	res := Asset{}

	// check permission
	err := u.Ctx.ValidatePermission("assets.detail")
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

// Get returns the list of Asset data.
func (u UseCaseHandler) Get() (app.ListModel, error) {
	res := app.ListModel{}

	// check permission
	err := u.Ctx.ValidatePermission("assets.list")
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
		err = app.Query().PaginationInfo(tx, &Asset{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	// return data count if $per_page set to 0
	if res.Results.PageContext.PerPage == 0 {
		return res, err
	}

	// find data
	data, err := app.Query().Find(tx, &Asset{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	res.SetData(data, u.Query)

	// save to cache and return if exists
	app.Cache().Set(cacheKey, res)
	return res, err
}

// Create creates a new data Asset with specified parameters.
func (u *UseCaseHandler) Create(p *ParamCreate) error {

	// check permission
	err := u.Ctx.ValidatePermission("assets.create")
	if err != nil {
		return err
	}

	// validate param
	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	// set default value for undefined field
	err = u.setDefaultValue(Asset{})
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

// UpdateByID updates the Asset data for the specified ID with specified parameters.
func (u UseCaseHandler) UpdateByID(id string, p *ParamUpdate) error {

	// check permission
	err := u.Ctx.ValidatePermission("assets.edit")
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

// PartiallyUpdateByID updates the Asset data for the specified ID with specified parameters.
func (u UseCaseHandler) PartiallyUpdateByID(id string, p *ParamPartiallyUpdate) error {

	// check permission
	err := u.Ctx.ValidatePermission("assets.edit")
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

// DeleteByID deletes the Asset data for the specified ID.
func (u UseCaseHandler) DeleteByID(id string, p *ParamDelete) error {

	// check permission
	err := u.Ctx.ValidatePermission("assets.delete")
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

// setDefaultValue set default value of undefined field when create or update Asset data.
func (u *UseCaseHandler) setDefaultValue(old Asset) error {

	if !old.ID.Valid {
		u.ID = app.NewNullUUID()
	} else {
		u.ID = old.ID
	}

	// validate category
	catKey := u.CategoryID.String
	if !u.CategoryID.Valid || u.CategoryID.String == "" {
		catKey = u.CategoryCode.String
	}
	if catKey != "" {
		cat, err := category.UseCaseHandler{Ctx: u.Ctx, Query: url.Values{}}.GetByID(catKey)
		if err != nil {
			return err
		}

		//update depreciation price based on category economic age
		if cat.Ages.Int64 > 0 && u.Price.Float64 > 0 {
			// Depresiasi per bulan = (Harga Asset - Salvage Amount) / Umur Ekonomis (bulan)
			depreciation := (u.Price.Float64 - u.SalvageAmount.Float64) / float64(cat.Ages.Int64)
			u.DepreciationAmountPerMonth.Set(depreciation)
		} else {
			u.DepreciationAmountPerMonth.Set(0)
		}
	}

	// validate attachment
	if u.AttachmentID.Valid && u.AttachmentID.String != "" {
		attUC := attachment.UseCase(*u.Ctx, url.Values{})
		att, err := attUC.GetByID(u.AttachmentID.String)
		if err != nil {
			return err
		}

		// Update data attachment
		upAtt := attachment.ParamUpdate{}
		upAtt.Endpoint.Set("assets")
		upAtt.DataId.Set(u.ID.String)
		err = attUC.UpdateByID(att.ID.String, &upAtt)
		if err != nil {
			return err
		}
	}

	// init input date
	if !u.InputDate.Valid && !old.InputDate.Valid {
		if u.CreatedAt.Valid {
			u.InputDate.Set(u.CreatedAt.Time)
		} else if old.CreatedAt.Valid {
			u.InputDate.Set(old.CreatedAt.Time)
		}
	}

	if !u.SalvageAmount.Valid && old.SalvageAmount.Valid {
		u.SalvageAmount = old.SalvageAmount
	} else if !old.SalvageAmount.Valid && !u.SalvageAmount.Valid {
		u.SalvageAmount.Set(0)
	}

	// Penentuan kode
	if u.Code.Valid && u.Code.String != "" {
		// Jika kode dikirim dan berbeda dengan data lama, cek ke DB
		if !old.Code.Valid || u.Code.String != old.Code.String {
			err := app.Common().IsFieldValueExists(u.Ctx, u.EndPoint(), "Code", u.TableName(), "code", u.Code.String)
			if err != nil {
				return err
			}
		}
		// Jika kode dikirim dan data lama tidak ada, cek ke DB (sudah tercakup di atas)
	} else {
		// Jika tidak kirim kode dan data lama ada, gunakan data lama
		if old.Code.Valid && old.Code.String != "" {
			u.Code = old.Code
		} else {
			// Jika tidak kirim kode dan data lama tidak ada, generate baru
			newCode, err := app.Common().GenerateCode(u.Ctx, u.TableName(), "code", u.Name.String)
			if err != nil {
				return err
			}
			u.Code.Set(newCode)
		}
	}

	//hitung nilai depresiasi & current value
	err := u.SetCurrentValue()
	if err != nil {
		return err
	}

	return nil
}

func (u *UseCaseHandler) SetCurrentValue() error {
	if !u.InputDate.Valid || u.Price.Float64 <= 0 || u.DepreciationAmountPerMonth.Float64 <= 0 {
		return nil
	}

	// Hitung jumlah bulan berlalu sejak InputDate
	start := u.InputDate.Time
	now := time.Now().UTC()
	months := int((now.Year()-start.Year())*12 + int(now.Month()) - int(start.Month()))
	if now.Day() < start.Day() {
		months--
	}
	if months < 0 {
		months = 0
	}

	// Hitung total depresiasi
	totalDepreciation := float64(months) * u.DepreciationAmountPerMonth.Float64

	// Hitung nilai ekonomis saat ini, minimal salvage amount
	currentValue := u.Price.Float64 - totalDepreciation
	if currentValue < u.SalvageAmount.Float64 {
		currentValue = u.SalvageAmount.Float64
	}

	// Set hasil ke struct (pastikan field CurrentValue & TotalDepreciation ada)
	u.CurrentValue.Set(currentValue)
	u.DepreciationAmount.Set(totalDepreciation)

	return nil
}

func JobUpdateAssetValue() {
	tx, err := app.DB().Conn("main")
	if err != nil {
		return
	}

	// get all assets
	assets := []Asset{}
	err = tx.Find(&assets).Error
	if err != nil {
		return
	}

	for _, u := range assets {
		if !u.SalvageAmount.Valid {
			u.SalvageAmount.Set(0)
		}

		if u.CategoryID.Valid {
			cat := category.Category{}
			err = tx.Model(&category.Category{}).Where("id = ?", u.CategoryID.String).First(&cat).Error
			if err != nil {
				return
			}

			//update depreciation price based on category economic age
			if cat.Ages.Int64 > 0 && u.Price.Float64 > 0 {
				// Depresiasi per bulan = (Harga Asset - Salvage Amount) / Umur Ekonomis (bulan)
				depreciation := (u.Price.Float64 - u.SalvageAmount.Float64) / float64(cat.Ages.Int64)
				u.DepreciationAmountPerMonth.Set(depreciation)
			} else {
				u.DepreciationAmountPerMonth.Set(0)
			}
		}

		// Hitung jumlah bulan berlalu sejak InputDate
		if !u.InputDate.Valid {
			u.InputDate.Set(u.CreatedAt.Time)
		}

		start := u.InputDate.Time
		now := time.Now().UTC()
		months := int((now.Year()-start.Year())*12 + int(now.Month()) - int(start.Month()))
		if now.Day() < start.Day() {
			months--
		}
		if months < 0 {
			months = 0
		}

		// Hitung total depresiasi
		totalDepreciation := float64(0)
		if u.DepreciationAmountPerMonth.Valid {
			totalDepreciation = float64(months) * u.DepreciationAmountPerMonth.Float64
		}

		// Hitung nilai ekonomis saat ini, minimal salvage amount
		currentValue := u.Price.Float64 - totalDepreciation
		if currentValue < u.SalvageAmount.Float64 {
			currentValue = u.SalvageAmount.Float64
		}

		// Set hasil ke struct (pastikan field CurrentValue & TotalDepreciation ada)
		u.CurrentValue.Set(currentValue)
		u.DepreciationAmount.Set(totalDepreciation)

		// save data to db
		err = tx.Model(&u).Where("id = ?", u.ID.String).Updates(&u).Error
		if err != nil {
			return
		}
	}

	app.Cache().DeleteWithPrefix(Asset{}.EndPoint())
}

func (u UseCaseHandler) GetDepreciation(id string) ([]DepreciationList, error) {
	res := []DepreciationList{}

	// get data asset
	asset, err := u.GetByID(id)
	if err != nil {
		return res, err
	}

	// generate depreciation list
	if !asset.InputDate.Valid || asset.Price.Float64 <= 0 || asset.DepreciationAmountPerMonth.Float64 <= 0 {
		return res, nil
	}

	// economic age in months
	economicAge := int64(0)
	if asset.CategoryEconomicAges.Valid {
		economicAge = asset.CategoryEconomicAges.Int64
	}
	if economicAge == 0 {
		return res, nil
	}

	initialAmount := asset.Price.Float64
	assetAmount := asset.Price.Float64
	salvageAmount := asset.SalvageAmount.Float64
	depPerMonth := asset.DepreciationAmountPerMonth.Float64

	for month := int64(1); month <= economicAge; month++ {
		if month != 1 {
			assetAmount -= depPerMonth
			if salvageAmount != 0 {
				if assetAmount < salvageAmount {
					assetAmount = salvageAmount
				}
			} else {
				if assetAmount < 0 {
					assetAmount = 0
				}
			}
		}
		date := asset.InputDate.Time.AddDate(0, int(month-1), 0)
		depreciation := depPerMonth
		economicAmount := assetAmount - depreciation
		if salvageAmount != 0 {
			if economicAmount < salvageAmount {
				economicAmount = salvageAmount
			}
		} else {
			if economicAmount < 0 {
				economicAmount = 0
			}
		}
		res = append(res, DepreciationList{
			Date:               app.NewNullDate(date),
			Month:              app.NewNullInt64(month),
			InitialAmount:      app.NewNullFloat64(initialAmount),
			AssetAmount:        app.NewNullFloat64(assetAmount),
			DepreciationAmount: app.NewNullFloat64(depreciation),
			EconomicAmount:     app.NewNullFloat64(economicAmount),
		})
	}

	return res, nil
}
