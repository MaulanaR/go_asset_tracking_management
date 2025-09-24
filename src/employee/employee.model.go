package employee

import "github.com/maulanar/go_asset_tracking_management/app"

// Employee is the main model of Employee data. It provides a convenient interface for app.ModelInterface
type Employee struct {
	app.Model
	ID                    app.NullUUID   `json:"id"                     db:"m.id"              gorm:"column:id;primaryKey"`
	Code                  app.NullString `json:"code"                   db:"m.code"            gorm:"column:code"`
	Name                  app.NullString `json:"name"                   db:"m.name"            gorm:"column:name"`
	DepartmentID          app.NullUUID   `json:"department.id"          db:"m.department_id"   gorm:"column:department_id"`
	DepartmentCode        app.NullString `json:"department.code"        db:"dpt.code"          gorm:"-"`
	DepartmentName        app.NullString `json:"department.name"        db:"dpt.name"          gorm:"-"`
	DepartmentDescription app.NullText   `json:"department.description" db:"dpt.description"   gorm:"-"`
	BranchID              app.NullUUID   `json:"branch.id"              db:"m.branch_id"       gorm:"column:branch_id"`
	BranchCode            app.NullString `json:"branch.code"            db:"brc.code"          gorm:"-"`
	BranchName            app.NullString `json:"branch.name"            db:"brc.name"          gorm:"-"`
	BranchAddress         app.NullText   `json:"branch.address"         db:"brc.address"       gorm:"-"`
	Address               app.NullText   `json:"address"                db:"m.address"         gorm:"column:address"`
	Phone                 app.NullString `json:"phone"                  db:"m.phone"           gorm:"column:phone"`
	AttachmentID          app.NullText   `json:"attachment.id"          db:"m.attachment_id"   gorm:"column:attachment_id"`
	AttachmentName        app.NullText   `json:"attachment.name"        db:"att.name"          gorm:"-"`
	AttachmentPath        app.NullText   `json:"attachment.path"        db:"att.path"          gorm:"-"`
	AttachmentURL         app.NullText   `json:"attachment.url"         db:"att.url"           gorm:"-"`
	Email                 app.NullString `json:"email"                  db:"m.email"           gorm:"column:email"`
	IsActive              app.NullBool   `json:"is_active"              db:"m.is_active"       gorm:"column:is_active;default:true"`

	CreatedAt app.NullDateTime `json:"created_at"             db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at"             db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at"             db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Employee end point, it used for cache key, etc.
func (Employee) EndPoint() string {
	return "employees"
}

// TableVersion returns the versions of the Employee table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Employee) TableVersion() string {
	return "25.09.241722"
}

// TableName returns the name of the Employee table in the database.
func (Employee) TableName() string {
	return "employees"
}

// TableAliasName returns the table alias name of the Employee table, used for querying.
func (Employee) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Employee data in the database, used for querying.
func (m *Employee) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "departments", "dpt", []map[string]any{{"column1": "dpt.id", "column2": "m.department_id"}})
	m.AddRelation("left", "branches", "brc", []map[string]any{{"column1": "brc.id", "column2": "m.branch_id"}})
	m.AddRelation("left", "attachments", "att", []map[string]any{{"column1": "att.id", "column2": "m.attachment_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Employee data in the database, used for querying.
func (m *Employee) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Employee data in the database, used for querying.
func (m *Employee) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Employee data in the database, used for querying.
func (m *Employee) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Employee schema, used for querying.
func (m *Employee) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Employee schema in the open api documentation.
func (Employee) OpenAPISchemaName() string {
	return "Employee"
}

// GetOpenAPISchema returns the Open API Schema of the Employee in the open api documentation.
func (m *Employee) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type EmployeeList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the EmployeeList schema in the open api documentation.
func (EmployeeList) OpenAPISchemaName() string {
	return "EmployeeList"
}

// GetOpenAPISchema returns the Open API Schema of the EmployeeList in the open api documentation.
func (p *EmployeeList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Employee{})
}

// ParamCreate is the expected parameters for create a new Employee data.
type ParamCreate struct {
	UseCaseHandler
	Name  app.NullString `json:"name"  db:"m.name"  gorm:"column:name"  validate:"required"`
	Phone app.NullString `json:"phone" db:"m.phone" gorm:"column:phone" validate:"required"`
}

// ParamUpdate is the expected parameters for update the Employee data.
type ParamUpdate struct {
	UseCaseHandler
	Name   app.NullString `json:"name"   db:"m.name"  gorm:"column:name"  validate:"required"`
	Phone  app.NullString `json:"phone"  db:"m.phone" gorm:"column:phone" validate:"required"`
	Reason app.NullString `json:"reason" gorm:"-"            validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the Employee data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the Employee data.
type ParamDelete struct {
	UseCaseHandler
}
