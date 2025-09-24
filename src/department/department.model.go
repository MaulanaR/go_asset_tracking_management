package department

import "github.com/maulanar/go_asset_tracking_management/app"

// Department is the main model of Department data. It provides a convenient interface for app.ModelInterface
type Department struct {
	app.Model
	ID          app.NullUUID   `json:"id"          db:"m.id"              gorm:"column:id;primaryKey"`
	Code        app.NullString `json:"code"        db:"m.code"            gorm:"column:code"`
	Name        app.NullString `json:"name"        db:"m.name"            gorm:"column:name"`
	Description app.NullText   `json:"description" db:"m.description"     gorm:"column:description"`
	IsActive    app.NullBool   `json:"is_active"   db:"m.is_active"       gorm:"column:is_active;default:true"`

	CreatedAt app.NullDateTime `json:"created_at"  db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at"  db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at"  db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Department end point, it used for cache key, etc.
func (Department) EndPoint() string {
	return "departments"
}

// TableVersion returns the versions of the Department table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Department) TableVersion() string {
	return "25.09.241152"
}

// TableName returns the name of the Department table in the database.
func (Department) TableName() string {
	return "departments"
}

// TableAliasName returns the table alias name of the Department table, used for querying.
func (Department) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Department data in the database, used for querying.
func (m *Department) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Department data in the database, used for querying.
func (m *Department) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Department data in the database, used for querying.
func (m *Department) GetSorts() []map[string]any {
	// m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Department data in the database, used for querying.
func (m *Department) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Department schema, used for querying.
func (m *Department) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Department schema in the open api documentation.
func (Department) OpenAPISchemaName() string {
	return "Department"
}

// GetOpenAPISchema returns the Open API Schema of the Department in the open api documentation.
func (m *Department) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type DepartmentList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the DepartmentList schema in the open api documentation.
func (DepartmentList) OpenAPISchemaName() string {
	return "DepartmentList"
}

// GetOpenAPISchema returns the Open API Schema of the DepartmentList in the open api documentation.
func (p *DepartmentList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Department{})
}

// ParamCreate is the expected parameters for create a new Department data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Department data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the Department data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the Department data.
type ParamDelete struct {
	UseCaseHandler
}
