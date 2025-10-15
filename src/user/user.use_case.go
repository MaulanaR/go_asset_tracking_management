package user

import (
	"encoding/json"
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

// ✅ PERBAIKAN: Fungsi Get dengan Raw Query
func (u UseCaseHandler) Get() (app.ListModel, error) {
	res := app.ListModel{}

	// permission
	if err := u.Ctx.ValidatePermission("users.list"); err != nil {
		return res, err
	}

	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// pagination info
	res.Results.PageContext.Count,
		res.Results.PageContext.Page,
		res.Results.PageContext.PerPage,
		res.Results.PageContext.PageCount,
		err = app.Query().PaginationInfo(tx, &User{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	if res.Results.PageContext.PerPage == 0 {
		return res, err
	}

	// Pastikan page/perPage valid
	page := res.Results.PageContext.Page
	perPage := res.Results.PageContext.PerPage
	if page <= 0 {
		page = 1
	}
	if perPage <= 0 {
		perPage = 10
	}

	// ✅ SOLUSI: Gunakan Raw Query dengan Scan manual
	rows, err := tx.Raw(`
		SELECT 
			m.id,
			m.email,
			m.full_name,
			m.phone,
			m.is_active,
			m.created_at,
			m.updated_at,
			m.deleted_at,
			m.role_id,
			rl.name AS role_name,
			rl.acl AS role_acl
		FROM users AS m
		LEFT JOIN roles AS rl ON rl.id = m.role_id
		ORDER BY m.created_at DESC
		LIMIT ? OFFSET ?
	`, perPage, (page-1)*perPage).Rows()

	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	// Scan ke slice of User
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.FullName,
			&user.Phone,
			&user.IsActive,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
			&user.RoleID,
			&user.RoleName,
			&user.RoleACL,
		)
		if err != nil {
			return res, app.Error().New(http.StatusInternalServerError, err.Error())
		}
		users = append(users, user)
	}

	// Convert []User -> []map[string]any via JSON marshal/unmarshal
	b, err := json.Marshal(users)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	var rowsData []map[string]any
	if err := json.Unmarshal(b, &rowsData); err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	if rowsData == nil {
		rowsData = make([]map[string]any, 0)
	}

	// Set data dan cache
	res.SetData(rowsData, u.Query)
	app.Cache().Set(u.EndPoint()+"?"+u.Query.Encode(), res)

	return res, nil
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

	// Cek apakah email sudah terdaftar
	var existing User
	if err := tx.Where("email = ?", p.Email).First(&existing).Error; err == nil {
		return app.Error().New(http.StatusBadRequest, "Email is already registered")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// Hash password
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

	if p.RoleID != "" {
		user.RoleID.Set(p.RoleID)
	}

	// Simpan ke DB
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

// ✅ PERBAIKAN: Profile juga perlu di-fix
func (u UseCaseHandler) Profile(userID string) (User, error) {
	res := User{}

	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// ✅ Gunakan Raw Query dengan JOIN
	rows, err := tx.Raw(`
		SELECT 
			m.id,
			m.email,
			m.full_name,
			m.phone,
			m.is_active,
			m.created_at,
			m.updated_at,
			m.deleted_at,
			m.role_id,
			rl.name AS role_name,
			rl.acl AS role_acl
		FROM users AS m
		LEFT JOIN roles AS rl ON rl.id = m.role_id
		WHERE m.id = ?
	`, userID).Rows()

	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(
			&res.ID,
			&res.Email,
			&res.FullName,
			&res.Phone,
			&res.IsActive,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
			&res.RoleID,
			&res.RoleName,
			&res.RoleACL,
		)
		if err != nil {
			return res, app.Error().New(http.StatusInternalServerError, err.Error())
		}
	} else {
		return res, app.Error().New(http.StatusNotFound, "user not found")
	}

	return res, nil
}
