package user

import (
	"encoding/json"

	"github.com/maulanar/go_asset_tracking_management/app"
)

// User is the main model of User data.
type User struct {
	app.Model
	ID        app.NullUUID     `json:"id" db:"m.id" gorm:"column:id;primaryKey"`
	Email     app.NullString   `json:"email" db:"m.email" gorm:"column:email;unique"`
	Password  app.NullString   `json:"-" db:"m.password" gorm:"column:password"`
	FullName  app.NullString   `json:"full_name" db:"m.full_name" gorm:"column:full_name"`
	Phone     app.NullString   `json:"phone" db:"m.phone" gorm:"column:phone"`
	IsActive  app.NullBool     `json:"is_active" db:"m.is_active" gorm:"column:is_active;default:true"`
	CreatedAt app.NullDateTime `json:"created_at" db:"m.created_at" gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at" db:"m.updated_at" gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at" db:"m.deleted_at" gorm:"column:deleted_at"`

	RoleID app.NullUUID `json:"-" db:"m.role_id" gorm:"column:role_id"`
	// ✅ PERBAIKAN: Hapus gorm:"-" agar bisa di-scan
	RoleName app.NullString `json:"-" db:"role_name" gorm:"-:all"` // atau gunakan gorm:"<-:false"
	RoleACL  app.NullJSON   `json:"-" db:"role_acl" gorm:"-:all"`  // atau gunakan gorm:"<-:false"
}

// MarshalJSON custom untuk membuat nested role object
func (u User) MarshalJSON() ([]byte, error) {
	type Alias User

	return json.Marshal(&struct {
		*Alias
		Role map[string]any `json:"role"`
	}{
		Alias: (*Alias)(&u),
		Role: map[string]any{
			"id":   u.RoleID,
			"name": u.RoleName,
			"acl":  u.RoleACL,
		},
	})
}

// EndPoint returns endpoint name
func (User) EndPoint() string {
	return "users"
}

func (User) TableVersion() string {
	return "25.10.151953"
}

func (User) TableName() string {
	return "users"
}

func (User) TableAliasName() string {
	return "m"
}

func (m *User) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "roles", "rl", []map[string]any{{"column1": "rl.id", "column2": "m.role_id"}})

	return m.Relations
}
func (m *User) GetFilters() []map[string]any { return m.Filters }
func (m *User) GetSorts() []map[string]any   { return m.Sorts }
func (m *User) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}
func (m *User) GetSchema() map[string]any { return m.SetSchema(m) }

func (User) OpenAPISchemaName() string           { return "User" }
func (m *User) GetOpenAPISchema() map[string]any { return m.SetOpenAPISchema(m) }

type UserList struct {
	app.ListModel
}

func (UserList) OpenAPISchemaName() string { return "UserList" }
func (p *UserList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&User{})
}

type ParamRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	RoleID   string `json:"role_id"`
}

type ParamLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
