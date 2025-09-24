package asset

import "github.com/maulanar/go_asset_tracking_management/app"

// Asset is the main model of Asset data. It provides a convenient interface for app.ModelInterface
type Asset struct {
	app.Model
	ID           app.NullUUID    `json:"id"            db:"m.id"              gorm:"column:id;primaryKey"`
	Code         app.NullString  `json:"code"          db:"m.code"            gorm:"column:code"`
	Name         app.NullString  `json:"name"          db:"m.name"            gorm:"column:name"`
	Price        app.NullFloat64 `json:"price"         db:"m.price"           gorm:"column:price"`
	Attachment   app.NullText    `json:"attachment"    db:"m.attachment"      gorm:"column:attachment"`
	CategoryID   app.NullUUID    `json:"category.id"   db:"m.category_id"     gorm:"column:category_id"`
	CategoryCode app.NullString  `json:"category.code" db:"c.code"            gorm:"-"`
	CategoryName app.NullString  `json:"category.name" db:"c.name"            gorm:"-"`
	Status       app.NullString  `json:"status"        db:"m.status"          gorm:"column:status"`

	CreatedAt app.NullDateTime `json:"created_at"    db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at"    db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at"    db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Asset end point, it used for cache key, etc.
func (Asset) EndPoint() string {
	return "assets"
}

// TableVersion returns the versions of the Asset table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Asset) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the Asset table in the database.
func (Asset) TableName() string {
	return "assets"
}

// TableAliasName returns the table alias name of the Asset table, used for querying.
func (Asset) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Asset data in the database, used for querying.
func (m *Asset) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "categories", "c", []map[string]any{{"column1": "c.id", "column2": "m.category_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Asset data in the database, used for querying.
func (m *Asset) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Asset data in the database, used for querying.
func (m *Asset) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Asset data in the database, used for querying.
func (m *Asset) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Asset schema, used for querying.
func (m *Asset) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Asset schema in the open api documentation.
func (Asset) OpenAPISchemaName() string {
	return "Asset"
}

// GetOpenAPISchema returns the Open API Schema of the Asset in the open api documentation.
func (m *Asset) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type AssetList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the AssetList schema in the open api documentation.
func (AssetList) OpenAPISchemaName() string {
	return "AssetList"
}

// GetOpenAPISchema returns the Open API Schema of the AssetList in the open api documentation.
func (p *AssetList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Asset{})
}

// ParamCreate is the expected parameters for create a new Asset data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Asset data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the Asset data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the Asset data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
