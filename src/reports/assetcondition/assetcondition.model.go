package assetcondition

import "github.com/maulanar/go_asset_tracking_management/app"

// AssetCondition is the main model of AssetCondition data. It provides a convenient interface for app.ModelInterface
type AssetCondition struct {
	app.Model
	ID app.NullUUID `json:"id"         db:"m.id"              gorm:"column:id;primaryKey"`

	CreatedAt app.NullDateTime `json:"created_at" db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at" db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at" db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the AssetCondition end point, it used for cache key, etc.
func (AssetCondition) EndPoint() string {
	return "asset_conditions"
}

// TableVersion returns the versions of the AssetCondition table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (AssetCondition) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the AssetCondition table in the database.
func (AssetCondition) TableName() string {
	return "asset_conditions"
}

// TableAliasName returns the table alias name of the AssetCondition table, used for querying.
func (AssetCondition) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the AssetCondition data in the database, used for querying.
func (m *AssetCondition) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the AssetCondition data in the database, used for querying.
func (m *AssetCondition) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the AssetCondition data in the database, used for querying.
func (m *AssetCondition) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the AssetCondition data in the database, used for querying.
func (m *AssetCondition) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the AssetCondition schema, used for querying.
func (m *AssetCondition) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the AssetCondition schema in the open api documentation.
func (AssetCondition) OpenAPISchemaName() string {
	return "AssetCondition"
}

// GetOpenAPISchema returns the Open API Schema of the AssetCondition in the open api documentation.
func (m *AssetCondition) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type AssetConditionList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the AssetConditionList schema in the open api documentation.
func (AssetConditionList) OpenAPISchemaName() string {
	return "AssetConditionList"
}

// GetOpenAPISchema returns the Open API Schema of the AssetConditionList in the open api documentation.
func (p *AssetConditionList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&AssetCondition{})
}

type Asset struct {
	app.Model
	ID             app.NullUUID    `json:"id"                       db:"m.id"                     gorm:"column:id;primaryKey"`
	Code           app.NullString  `json:"code"                     db:"m.code"                   gorm:"column:code"`
	Name           app.NullString  `json:"name"                     db:"m.name"                   gorm:"column:name"`
	Price          app.NullFloat64 `json:"price"                    db:"m.price"                  gorm:"column:price"`
	AttachmentID   app.NullUUID    `json:"attachment.id"            db:"m.attachment_id"          gorm:"column:attachment_id"`
	AttachmentName app.NullText    `json:"attachment.name"          db:"att.name"                 gorm:"-"`
	AttachmentPath app.NullText    `json:"attachment.path"          db:"att.path"                 gorm:"-"`
	AttachmentURL  app.NullText    `json:"attachment.url"           db:"att.url"                  gorm:"-"`

	CategoryID          app.NullUUID   `json:"category.id"              db:"m.category_id"            gorm:"column:category_id"`
	CategoryCode        app.NullString `json:"category.code"            db:"cat.code"                 gorm:"-"`
	CategoryName        app.NullString `json:"category.name"            db:"cat.name"                 gorm:"-"`
	CategoryDescription app.NullText   `json:"category.description"     db:"cat.description"          gorm:"-"`

	ConditionID          app.NullUUID   `json:"condition.id"             db:"emp_ass.condition_id"     gorm:"-"`
	ConditionCode        app.NullString `json:"condition.code"           db:"emp_ass_cond.code"        gorm:"-"`
	ConditionName        app.NullString `json:"condition.name"           db:"emp_ass_cond.name"        gorm:"-"`
	ConditionDescription app.NullText   `json:"condition.description"    db:"emp_ass_cond.description" gorm:"-"`
	ConditionColor       app.NullString `json:"condition.color"          db:"emp_ass_cond.color"       gorm:"-"`

	DepartmentID          app.NullUUID   `json:"department.id"            db:"emp_ass_dpt.id"           gorm:"-"`
	DepartmentCode        app.NullString `json:"department.code"          db:"emp_ass_dpt.code"         gorm:"-"`
	DepartmentName        app.NullString `json:"department.name"          db:"emp_ass_dpt.name"         gorm:"-"`
	DepartmentDescription app.NullText   `json:"department.description"   db:"emp_ass_dpt.description"  gorm:"-"`

	JobPositionID          app.NullUUID   `json:"job_position.id"          db:"emp_job_pos.id"           gorm:"-"`
	JobPositionCode        app.NullString `json:"job_position.code"        db:"emp_job_pos.code"         gorm:"-"`
	JobPositionName        app.NullString `json:"job_position.name"        db:"emp_job_pos.name"         gorm:"-"`
	JobPositionDescription app.NullText   `json:"job_position.description" db:"emp_job_pos.description"  gorm:"-"`

	EmployeeID       app.NullUUID   `json:"employee.id"              db:"emp.id"                   gorm:"-"`
	EmployeeCode     app.NullString `json:"employee.code"            db:"emp.code"                 gorm:"-"`
	EmployeeName     app.NullString `json:"employee.name"            db:"emp.name"                 gorm:"-"`
	EmployeeAddress  app.NullText   `json:"employee.address"         db:"emp.address"              gorm:"-"`
	EmployeePhone    app.NullString `json:"employee.phone"           db:"emp.phone"                gorm:"-"`
	EmployeeEmail    app.NullString `json:"employee.email"           db:"emp.email"                gorm:"-"`
	EmployeeIsActive app.NullBool   `json:"employee.is_active"       db:"emp.is_active"            gorm:"-"`

	BranchID      app.NullUUID   `json:"branch.id"                db:"emp_ass_brc.id"           gorm:"-"`
	BranchCode    app.NullString `json:"branch.code"              db:"emp_ass_brc.code"         gorm:"-"`
	BranchName    app.NullString `json:"branch.name"              db:"emp_ass_brc.name"         gorm:"-"`
	BranchAddress app.NullText   `json:"branch.address"           db:"emp_ass_brc.address"      gorm:"-"`

	Status    app.NullString   `json:"status"                   db:"m.status"                 gorm:"column:status"        validate:"omitempty,oneof=available unavailable"`
	CreatedAt app.NullDateTime `json:"created_at"               db:"m.created_at"             gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at"               db:"m.updated_at"             gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at"               db:"m.deleted_at,hide"        gorm:"column:deleted_at"`
}

// EndPoint returns the Asset end point, it used for cache key, etc.
func (Asset) EndPoint() string {
	return "assets"
}

// TableVersion returns the versions of the Asset table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Asset) TableVersion() string {
	return "25.09.242030"
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
	m.AddRelation("left", "categories", "cat", []map[string]any{{"column1": "cat.id", "column2": "m.category_id"}})
	m.AddRelation("left", "departments", "dep", []map[string]any{{"column1": "dep.id", "column2": "m.department_id"}})
	m.AddRelation("left", "attachments", "att", []map[string]any{{"column1": "att.id", "column2": "m.attachment_id"}})

	// search to employee_assets
	m.AddRelation("left", `(
  SELECT DISTINCT ON (ea.asset_id)
         ea.date,
         ea.asset_id,
         ea.condition_id,
         ea.employee_id
  FROM employee_assets ea
  WHERE ea.deleted_at IS NULL
  ORDER BY ea.asset_id, ea.date DESC, ea.id DESC
)`, "emp_ass", []map[string]any{{"column1": "emp_ass.asset_id", "column2": "m.id"}})
	m.AddRelation("left", "conditions", "emp_ass_cond", []map[string]any{{"column1": "emp_ass_cond.id", "column2": "emp_ass.condition_id"}})
	m.AddRelation("left", "employees", "emp", []map[string]any{{"column1": "emp.id", "column2": "emp_ass.employee_id"}})
	m.AddRelation("left", "departments", "emp_ass_dpt", []map[string]any{{"column1": "emp_ass_dpt.id", "column2": "emp.department_id"}})
	m.AddRelation("left", "branches", "emp_ass_brc", []map[string]any{{"column1": "emp_ass_brc.id", "column2": "emp.branch_id"}})
	m.AddRelation("left", "job_positions", "emp_job_pos", []map[string]any{{"column1": "emp_job_pos.id", "column2": "emp.job_position_id"}})

	return m.Relations
}

// GetFilters returns the filter of the Asset data in the database, used for querying.
func (m *Asset) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Asset data in the database, used for querying.
func (m *Asset) GetSorts() []map[string]any {
	// m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
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

type ViewData struct {
	CreatedAt string
	Datas     []Asset
}
