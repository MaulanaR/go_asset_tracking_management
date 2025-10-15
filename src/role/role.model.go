package role

import "github.com/maulanar/go_asset_tracking_management/app"

// Role is the main model of Role data.
type Role struct {
	app.Model
	ID       app.NullUUID   `json:"id"       db:"m.id"       gorm:"column:id;primaryKey"`
	Name     app.NullString `json:"name"    db:"m.name"    gorm:"column:name;uniqueIndex;not null" validate:"required"`
	ACL      app.NullJSON   `json:"acl" db:"m.acl" gorm:"column:acl;type:jsonb"`
	IsActive app.NullBool   `json:"is_active"   db:"m.is_active"       gorm:"column:is_active;default:true"`

	CreatedAt app.NullDateTime `json:"created_at"  db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at"  db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at"  db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Role end point, it used for cache key, etc.
func (Role) EndPoint() string {
	return "roles"
}

// TableVersion returns the versions of the Role table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Role) TableVersion() string {
	return "25.09.241152"
}

// TableName returns the name of the Role table in the database.
func (Role) TableName() string {
	return "roles"
}

// TableAliasName returns the table alias name of the Role table, used for querying.
func (Role) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Role data in the database, used for querying.
func (m *Role) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Role data in the database, used for querying.
func (m *Role) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Role data in the database, used for querying.
func (m *Role) GetSorts() []map[string]any {
	// m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Role data in the database, used for querying.
func (m *Role) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Role schema, used for querying.
func (m *Role) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Role schema in the open api documentation.
func (Role) OpenAPISchemaName() string {
	return "Role"
}

// GetOpenAPISchema returns the Open API Schema of the Role in the open api documentation.
func (m *Role) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type RoleList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the RoleList schema in the open api documentation.
func (RoleList) OpenAPISchemaName() string {
	return "RoleList"
}

// GetOpenAPISchema returns the Open API Schema of the RoleList in the open api documentation.
func (p *RoleList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Role{})
}

// ParamCreate is the expected parameters for create a new Role data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Role data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the Role data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the Role data.
type ParamDelete struct {
	UseCaseHandler
}
