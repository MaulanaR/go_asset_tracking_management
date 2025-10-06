package user

import (
	"errors"
	"net/http"
	"net/url"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/golang-jwt/jwt/v5"
	"github.com/maulanar/go_asset_tracking_management/app"
)

func UseCase(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	u := UseCaseHandler{Ctx: &ctx, Query: url.Values{}}
	if len(query) > 0 {
		u.Query = query[0]
	}
	return u
}

type UseCaseHandler struct {
	User
	Ctx   *app.Ctx   `json:"-" db:"-" gorm:"-"`
	Query url.Values `json:"-" db:"-" gorm:"-"`
}

func (u UseCaseHandler) Get() (app.ListModel, error) {
	res := app.ListModel{}

	// check permission
	err := u.Ctx.ValidatePermission("users.list")
	if err != nil {
		return res, err
	}
	// get from cache and return if exists
	cacheKey := u.EndPoint() + "?" + u.Query.Encode()
	// err = app.Cache().Get(cacheKey, &res)
	// if err == nil {
	// 	return res, err
	// }

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
		err = app.Query().PaginationInfo(tx, &User{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	// return data count if $per_page set to 0
	if res.Results.PageContext.PerPage == 0 {
		return res, err
	}

	// find data
	data, err := app.Query().Find(tx, &User{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	res.SetData(data, u.Query)

	// save to cache and return if exists
	app.Cache().Set(cacheKey, res)
	return res, err
}

func (u UseCaseHandler) DeleteByID(id string) error {
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(500, err.Error())
	}

	err = tx.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return app.Error().New(500, err.Error())
	}

	return nil
}

// Register user
func (u *UseCaseHandler) Register(p *ParamRegister) error {
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// 🔹 Cek apakah email sudah terdaftar
	var existing User
	if err := tx.Where("email = ?", p.Email).First(&existing).Error; err == nil {
		return app.Error().New(http.StatusBadRequest, "Email is already registered")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// 🔹 Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(p.Password), 12)
	if err != nil {
		return err
	}

	user := User{}
	user.ID = app.NewNullUUID()
	user.Email.Set(p.Email)
	user.FullName.Set(p.FullName)
	user.Phone.Set(p.Phone)
	user.Password.Set(string(hash))
	user.IsActive.Set(true)

	// 🔹 Simpan ke DB
	if err := tx.Model(&User{}).Create(&user).Error; err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

// Login and return JWT token
func (u *UseCaseHandler) Login(p *ParamLogin) (string, error) {
	tx, err := u.Ctx.DB()
	invalidCreds := app.Error().New(http.StatusUnauthorized, "Invalid email or password")

	if err != nil {
		return "", err
	}

	user := User{}
	if err := tx.Where("email = ?", p.Email).First(&user).Error; err != nil {
		return "", invalidCreds
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(p.Password)) != nil {
		time.Sleep(500 * time.Millisecond)
		return "", invalidCreds
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID.String,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u UseCaseHandler) Profile(userID string) (User, error) {
	res := User{}

	// check permission kalau mau (opsional)
	// err := u.Ctx.ValidatePermission("users.detail")
	// if err != nil {
	// 	return res, err
	// }

	// prepare db
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// ambil user by ID
	err = tx.Where("id = ?", userID).First(&res).Error
	if err != nil {
		return res, app.Error().New(http.StatusNotFound, "user not found")
	}

	return res, nil
}
