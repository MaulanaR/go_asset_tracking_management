package condition

import "github.com/maulanar/go_asset_tracking_management/app"

// Condition is the main model of Condition data. It provides a convenient interface for app.ModelInterface
type Condition struct {
	app.Model
	ID          app.NullUUID   `json:"id"          db:"m.id"              gorm:"column:id;primaryKey"`
	Code        app.NullString `json:"code"        db:"m.code"            gorm:"column:code"`
	Name        app.NullString `json:"name"        db:"m.name"            gorm:"column:name"`
	Color       app.NullString `json:"color"       db:"m.color"           gorm:"column:color"`
	Description app.NullText   `json:"description" db:"m.description"     gorm:"column:description"`
	IsActive    app.NullBool   `json:"is_active"   db:"m.is_active"       gorm:"column:is_active;default:true"`

	CreatedAt app.NullDateTime `json:"created_at"  db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at"  db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at"  db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Condition end point, it used for cache key, etc.
func (Condition) EndPoint() string {
	return "conditions"
}

// TableVersion returns the versions of the Condition table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Condition) TableVersion() string {
	return "25.09.261700"
}

// TableName returns the name of the Condition table in the database.
func (Condition) TableName() string {
	return "conditions"
}

// TableAliasName returns the table alias name of the Condition table, used for querying.
func (Condition) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Condition data in the database, used for querying.
func (m *Condition) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Condition data in the database, used for querying.
func (m *Condition) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Condition data in the database, used for querying.
func (m *Condition) GetSorts() []map[string]any {
	// m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Condition data in the database, used for querying.
func (m *Condition) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Condition schema, used for querying.
func (m *Condition) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Condition schema in the open api documentation.
func (Condition) OpenAPISchemaName() string {
	return "Condition"
}

// GetOpenAPISchema returns the Open API Schema of the Condition in the open api documentation.
func (m *Condition) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type ConditionList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the ConditionList schema in the open api documentation.
func (ConditionList) OpenAPISchemaName() string {
	return "ConditionList"
}

// GetOpenAPISchema returns the Open API Schema of the ConditionList in the open api documentation.
func (p *ConditionList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Condition{})
}

// ParamCreate is the expected parameters for create a new Condition data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Condition data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the Condition data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the Condition data.
type ParamDelete struct {
	UseCaseHandler
}
