package branch

import "github.com/maulanar/go_asset_tracking_management/app"

// Branch is the main model of Branch data. It provides a convenient interface for app.ModelInterface
type Branch struct {
	app.Model
	ID       app.NullUUID   `json:"id"         db:"m.id"              gorm:"column:id;primaryKey"`
	Code     app.NullString `json:"code"       db:"m.code"            gorm:"column:code"`
	Name     app.NullString `json:"name"       db:"m.name"            gorm:"column:name"`
	Address  app.NullText   `json:"address"    db:"m.address"         gorm:"column:address"`
	IsActive app.NullBool   `json:"is_active"  db:"m.is_active"       gorm:"column:is_active"`

	CreatedAt app.NullDateTime `json:"created_at" db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at" db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at" db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Branch end point, it used for cache key, etc.
func (Branch) EndPoint() string {
	return "branches"
}

// TableVersion returns the versions of the Branch table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Branch) TableVersion() string {
	return "25.09.241535"
}

// TableName returns the name of the Branch table in the database.
func (Branch) TableName() string {
	return "branches"
}

// TableAliasName returns the table alias name of the Branch table, used for querying.
func (Branch) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Branch data in the database, used for querying.
func (m *Branch) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Branch data in the database, used for querying.
func (m *Branch) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Branch data in the database, used for querying.
func (m *Branch) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Branch data in the database, used for querying.
func (m *Branch) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Branch schema, used for querying.
func (m *Branch) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Branch schema in the open api documentation.
func (Branch) OpenAPISchemaName() string {
	return "Branch"
}

// GetOpenAPISchema returns the Open API Schema of the Branch in the open api documentation.
func (m *Branch) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type BranchList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the BranchList schema in the open api documentation.
func (BranchList) OpenAPISchemaName() string {
	return "BranchList"
}

// GetOpenAPISchema returns the Open API Schema of the BranchList in the open api documentation.
func (p *BranchList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Branch{})
}

// ParamCreate is the expected parameters for create a new Branch data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Branch data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the Branch data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the Branch data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
