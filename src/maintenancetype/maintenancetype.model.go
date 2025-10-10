package maintenancetype

import "github.com/maulanar/go_asset_tracking_management/app"

// MaintenanceType is the main model of MaintenanceType data. It provides a convenient interface for app.ModelInterface
type MaintenanceType struct {
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

// EndPoint returns the MaintenanceType end point, it used for cache key, etc.
func (MaintenanceType) EndPoint() string {
	return "maintenance_types"
}

// TableVersion returns the versions of the MaintenanceType table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (MaintenanceType) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the MaintenanceType table in the database.
func (MaintenanceType) TableName() string {
	return "maintenance_types"
}

// TableAliasName returns the table alias name of the MaintenanceType table, used for querying.
func (MaintenanceType) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the MaintenanceType data in the database, used for querying.
func (m *MaintenanceType) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the MaintenanceType data in the database, used for querying.
func (m *MaintenanceType) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the MaintenanceType data in the database, used for querying.
func (m *MaintenanceType) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the MaintenanceType data in the database, used for querying.
func (m *MaintenanceType) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the MaintenanceType schema, used for querying.
func (m *MaintenanceType) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the MaintenanceType schema in the open api documentation.
func (MaintenanceType) OpenAPISchemaName() string {
	return "MaintenanceType"
}

// GetOpenAPISchema returns the Open API Schema of the MaintenanceType in the open api documentation.
func (m *MaintenanceType) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type MaintenanceTypeList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the MaintenanceTypeList schema in the open api documentation.
func (MaintenanceTypeList) OpenAPISchemaName() string {
	return "MaintenanceTypeList"
}

// GetOpenAPISchema returns the Open API Schema of the MaintenanceTypeList in the open api documentation.
func (p *MaintenanceTypeList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&MaintenanceType{})
}

// ParamCreate is the expected parameters for create a new MaintenanceType data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the MaintenanceType data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the MaintenanceType data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the MaintenanceType data.
type ParamDelete struct {
	UseCaseHandler
}
