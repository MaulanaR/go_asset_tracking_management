package attachment

import (
	"mime/multipart"

	"github.com/maulanar/go_asset_tracking_management/app"
)

// Attachment is the main model of Attachment data. It provides a convenient interface for app.ModelInterface
type Attachment struct {
	app.Model
	ID        app.NullUUID          `json:"id"         db:"m.id"              gorm:"column:id;primaryKey"`
	File      *multipart.FileHeader `json:"-"          form:"file" db:"-"                 gorm:"-"`
	Endpoint  app.NullString        `json:"endpoint"   db:"m.endpoint"        gorm:"column:endpoint"`
	DataId    app.NullUUID          `json:"data_id"    db:"m.data_id"         gorm:"column:data_id"`
	Name      app.NullText          `json:"name"       db:"m.name"            gorm:"column:name"`
	Path      app.NullText          `json:"path"       db:"m.path"            gorm:"column:path"`
	Url       app.NullText          `json:"url"        db:"m.url"             gorm:"column:url"`
	Extension app.NullText          `json:"extension"  db:"m.extension"       gorm:"column:extension"`

	CreatedAt app.NullDateTime `json:"created_at" db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at" db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at" db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Attachment end point, it used for cache key, etc.
func (Attachment) EndPoint() string {
	return "attachments"
}

// TableVersion returns the versions of the Attachment table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Attachment) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the Attachment table in the database.
func (Attachment) TableName() string {
	return "attachments"
}

// TableAliasName returns the table alias name of the Attachment table, used for querying.
func (Attachment) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Attachment data in the database, used for querying.
func (m *Attachment) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Attachment data in the database, used for querying.
func (m *Attachment) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Attachment data in the database, used for querying.
func (m *Attachment) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Attachment data in the database, used for querying.
func (m *Attachment) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Attachment schema, used for querying.
func (m *Attachment) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Attachment schema in the open api documentation.
func (Attachment) OpenAPISchemaName() string {
	return "Attachment"
}

// GetOpenAPISchema returns the Open API Schema of the Attachment in the open api documentation.
func (m *Attachment) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type AttachmentList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the AttachmentList schema in the open api documentation.
func (AttachmentList) OpenAPISchemaName() string {
	return "AttachmentList"
}

// GetOpenAPISchema returns the Open API Schema of the AttachmentList in the open api documentation.
func (p *AttachmentList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Attachment{})
}

// ParamCreate is the expected parameters for create a new Attachment data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Attachment data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the Attachment data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the Attachment data.
type ParamDelete struct {
	UseCaseHandler
}
