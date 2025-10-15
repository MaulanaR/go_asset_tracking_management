package user

import (
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/maulanar/go_asset_tracking_management/app"
)

func REST() *RESTAPIHandler {
	return &RESTAPIHandler{}
}

type RESTAPIHandler struct {
	UseCase UseCaseHandler
}

func (r *RESTAPIHandler) injectDeps(c *fiber.Ctx) error {
	ctx, ok := c.Locals(app.CtxKey).(*app.Ctx)
	if !ok {
		return app.Error().New(http.StatusInternalServerError, "ctx not found")
	}
	ctx.FiberCtx = c
	r.UseCase = UseCase(*ctx, app.Query().Parse(c.OriginalURL()))
	return nil
}

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

func (r *RESTAPIHandler) Register(c *fiber.Ctx) error {
	if err := r.injectDeps(c); err != nil {
		return app.Error().Handler(c, err)
	}

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		FullName string `json:"full_name"`
		Phone    string `json:"phone"`
		RoleID   string `json:"role_id"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "Failed",
			"error":   "invalid request body",
			"message": "email, password, full_name, and phone are required",
		})
	}

	if body.Email == "" || body.Password == "" || body.FullName == "" || body.Phone == "" {
		resp := app.ListSingleModel{Ctx: r.UseCase.Ctx}
		resp.Status = "Failed"
		resp.Message = "email, password, full_name, and phone are required"
		resp.TimeStamp = time.Now()

		return c.Status(http.StatusBadRequest).JSON(resp)
	}

	if body.RoleID != "" {
		if _, err := uuid.Parse(body.RoleID); err != nil {
			resp := app.ListSingleModel{Ctx: r.UseCase.Ctx}
			resp.Status = "Failed"
			resp.Message = "invalid role_id format, must be valid UUID"
			resp.TimeStamp = time.Now()

			return c.Status(http.StatusBadRequest).JSON(resp)
		}
	}

	if err := r.UseCase.Register(&ParamRegister{
		Email:    body.Email,
		Password: body.Password,
		FullName: body.FullName,
		Phone:    body.Phone,
		RoleID:   body.RoleID,
	}); err != nil {
		return app.Error().Handler(c, err)
	}

	// ✅ SOLUSI A: Gunakan Raw Query
	tx, _ := r.UseCase.Ctx.DB()

	// Query raw SQL
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
		WHERE m.email = ?
	`, body.Email).Rows()

	if err != nil {
		return app.Error().Handler(c, app.Error().New(http.StatusInternalServerError, err.Error()))
	}
	defer rows.Close()

	var user User
	if rows.Next() {
		// Scan manual
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
			return app.Error().Handler(c, app.Error().New(http.StatusInternalServerError, err.Error()))
		}
	}

	resp := app.ListSingleModel{Ctx: r.UseCase.Ctx}
	resp.SetData(user, r.UseCase.Query)

	if r.UseCase.IsFlat() {
		return c.Status(http.StatusCreated).JSON(resp)
	}
	return c.Status(http.StatusCreated).JSON(app.NewJSON(resp).ToStructured().Data)
}

func (r *RESTAPIHandler) Login(c *fiber.Ctx) error {
	if err := r.injectDeps(c); err != nil {
		return app.Error().Handler(c, err)
	}

	var p ParamLogin
	if err := c.BodyParser(&p); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	token, err := r.UseCase.Login(&p)
	if err != nil {
		return app.Error().Handler(c, err)
	}

	resp := app.ListSingleModel{Ctx: r.UseCase.Ctx}
	resp.SetData(map[string]any{
		"access_token": token,
	}, r.UseCase.Query) // <- use r.UseCase.Query

	if r.UseCase.IsFlat() {
		return c.JSON(resp)
	}
	return c.JSON(app.NewJSON(resp).ToStructured().Data)
}

func (r *RESTAPIHandler) Profile(c *fiber.Ctx) error {
	if err := r.injectDeps(c); err != nil {
		return app.Error().Handler(c, err)
	}

	// ambil token dari Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
	}
	tokenStr := authHeader[len("Bearer "):]

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret"
	}

	// parse JWT
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token claims"})
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid subject"})
	}

	// ambil profil user dari DB
	profile, err := r.UseCase.Profile(userID)
	if err != nil {
		return app.Error().Handler(c, err)
	}

	return c.JSON(profile)
}

func (r *RESTAPIHandler) DeleteByID(c *fiber.Ctx) error {
	if err := r.injectDeps(c); err != nil {
		return app.Error().Handler(c, err)
	}

	id := c.Params("id")

	err := r.UseCase.DeleteByID(id)
	if err != nil {
		return app.Error().Handler(c, err)
	}

	return c.JSON(fiber.Map{"message": "user deleted successfully"})
}
