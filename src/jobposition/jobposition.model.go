package jobposition

import "github.com/maulanar/go_asset_tracking_management/app"

// JobPosition is the main model of JobPosition data. It provides a convenient interface for app.ModelInterface
type JobPosition struct {
	app.Model
	ID          app.NullUUID   `json:"id"          db:"m.id"              gorm:"column:id;primaryKey"`
	Code        app.NullString `json:"code"        db:"m.code"            gorm:"column:code"`
	Name        app.NullString `json:"name"        db:"m.name"            gorm:"column:name"`
	Description app.NullText   `json:"description" db:"m.description"     gorm:"column:description"`

	CreatedAt app.NullDateTime `json:"created_at"  db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at"  db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at"  db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the JobPosition end point, it used for cache key, etc.
func (JobPosition) EndPoint() string {
	return "job_positions"
}

// TableVersion returns the versions of the JobPosition table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (JobPosition) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the JobPosition table in the database.
func (JobPosition) TableName() string {
	return "job_positions"
}

// TableAliasName returns the table alias name of the JobPosition table, used for querying.
func (JobPosition) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the JobPosition data in the database, used for querying.
func (m *JobPosition) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the JobPosition data in the database, used for querying.
func (m *JobPosition) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the JobPosition data in the database, used for querying.
func (m *JobPosition) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the JobPosition data in the database, used for querying.
func (m *JobPosition) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the JobPosition schema, used for querying.
func (m *JobPosition) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the JobPosition schema in the open api documentation.
func (JobPosition) OpenAPISchemaName() string {
	return "JobPosition"
}

// GetOpenAPISchema returns the Open API Schema of the JobPosition in the open api documentation.
func (m *JobPosition) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type JobPositionList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the JobPositionList schema in the open api documentation.
func (JobPositionList) OpenAPISchemaName() string {
	return "JobPositionList"
}

// GetOpenAPISchema returns the Open API Schema of the JobPositionList in the open api documentation.
func (p *JobPositionList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&JobPosition{})
}

// ParamCreate is the expected parameters for create a new JobPosition data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the JobPosition data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the JobPosition data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the JobPosition data.
type ParamDelete struct {
	UseCaseHandler
}
