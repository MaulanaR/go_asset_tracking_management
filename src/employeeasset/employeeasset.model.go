package employeeasset

import "github.com/maulanar/go_asset_tracking_management/app"

// EmployeeAsset is the main model of EmployeeAsset data. It provides a convenient interface for app.ModelInterface
type EmployeeAsset struct {
	app.Model
	ID         app.NullUUID `json:"id"                              db:"m.id"                 gorm:"column:id;primaryKey"`
	Date       app.NullDate `json:"date"                            db:"m.date"               gorm:"column:date"`
	AssignDate app.NullDate `json:"assign_date"                     db:"m.assign_date"        gorm:"column:assign_date"`

	AssetID                  app.NullUUID    `json:"asset.id"                        db:"m.asset_id"           gorm:"column:asset_id"`
	AssetCode                app.NullString  `json:"asset.code"                      db:"ass.code"             gorm:"-"`
	AssetName                app.NullString  `json:"asset.name"                      db:"ass.name"             gorm:"-"`
	AssetPrice               app.NullFloat64 `json:"asset.price"                     db:"ass.price"            gorm:"-"`
	AssetAttachmentID        app.NullUUID    `json:"asset.attachment.id"             db:"ass.attachment_id"    gorm:"-"`
	AssetAttachmentName      app.NullText    `json:"asset.attachment.name"           db:"ass_att.name"         gorm:"-"`
	AssetAttachmentPath      app.NullText    `json:"asset.attachment.path"           db:"ass_att.path"         gorm:"-"`
	AssetAttachmentURL       app.NullText    `json:"asset.attachment.url"            db:"ass_att.url"          gorm:"-"`
	AssetCategoryID          app.NullUUID    `json:"asset.category.id"               db:"ass.category_id"      gorm:"-"`
	AssetCategoryCode        app.NullString  `json:"asset.category.code"             db:"ass_cat.code"         gorm:"-"`
	AssetCategoryName        app.NullString  `json:"asset.category.name"             db:"ass_cat.name"         gorm:"-"`
	AssetCategoryDescription app.NullText    `json:"asset.category.description"      db:"ass_cat.description"  gorm:"-"`
	AssetStatus              app.NullString  `json:"asset.status"                    db:"ass.status"           gorm:"-"`

	EmployeeID                    app.NullUUID   `json:"employee.id"                     db:"m.employee_id"        gorm:"column:employee_id"`
	EmployeeCode                  app.NullString `json:"employee.code"                   db:"emp.code"             gorm:"-"`
	EmployeeName                  app.NullString `json:"employee.name"                   db:"emp.name"             gorm:"-"`
	EmployeeDepartmentID          app.NullUUID   `json:"employee.department.id"          db:"emp.department_id"    gorm:"-"`
	EmployeeDepartmentCode        app.NullString `json:"employee.department.code"        db:"emp_dept.code"        gorm:"-"`
	EmployeeDepartmentName        app.NullString `json:"employee.department.name"        db:"emp_dept.name"        gorm:"-"`
	EmployeeDepartmentDescription app.NullText   `json:"employee.department.description" db:"emp_dept.description" gorm:"-"`
	EmployeeBranchID              app.NullUUID   `json:"employee.branch.id"              db:"emp.branch_id"        gorm:"-"`
	EmployeeBranchCode            app.NullString `json:"employee.branch.code"            db:"emp_brc.code"         gorm:"-"`
	EmployeeBranchName            app.NullString `json:"employee.branch.name"            db:"emp_brc.name"         gorm:"-"`
	EmployeeBranchAddress         app.NullText   `json:"employee.branch.address"         db:"emp_brc.address"      gorm:"-"`
	EmployeeAddress               app.NullText   `json:"employee.address"                db:"emp.address"          gorm:"-"`
	EmployeePhone                 app.NullString `json:"employee.phone"                  db:"emp.phone"            gorm:"-"`
	EmployeeAttachmentID          app.NullUUID   `json:"employee.attachment.id"          db:"emp_att.id"           gorm:"-"`
	EmployeeAttachmentName        app.NullText   `json:"employee.attachment.name"        db:"emp_att.name"         gorm:"-"`
	EmployeeAttachmentPath        app.NullText   `json:"employee.attachment.path"        db:"emp_att.path"         gorm:"-"`
	EmployeeAttachmentURL         app.NullText   `json:"employee.attachment.url"         db:"emp_att.url"          gorm:"-"`
	EmployeeEmail                 app.NullString `json:"employee.email"                  db:"emp.email"            gorm:"-"`
	EmployeeIsActive              app.NullBool   `json:"employee.is_active"              db:"emp.is_active"        gorm:"-"`

	ConditionID          app.NullUUID   `json:"condition.id"                    db:"m.condition_id"       gorm:"column:condition_id"`
	ConditionCode        app.NullString `json:"condition.code"                  db:"cond.code"            gorm:"-"`
	ConditionName        app.NullString `json:"condition.name"                  db:"cond.name"            gorm:"-"`
	ConditionDescription app.NullText   `json:"condition.description"           db:"cond.description"     gorm:"-"`

	AttachmentID   app.NullUUID `json:"attachment.id"                   db:"m.attachment_id"      gorm:"column:attachment_id"`
	AttachmentName app.NullText `json:"attachment.name"                 db:"att.name"             gorm:"-"`
	AttachmentPath app.NullText `json:"attachment.path"                 db:"att.path"             gorm:"-"`
	AttachmentURL  app.NullText `json:"attachment.url"                  db:"att.url"              gorm:"-"`

	CreatedAt app.NullDateTime `json:"created_at"                      db:"m.created_at"         gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at"                      db:"m.updated_at"         gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at"                      db:"m.deleted_at,hide"    gorm:"column:deleted_at"`
}

// EndPoint returns the EmployeeAsset end point, it used for cache key, etc.
func (EmployeeAsset) EndPoint() string {
	return "employee_assets"
}

// TableVersion returns the versions of the EmployeeAsset table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (EmployeeAsset) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the EmployeeAsset table in the database.
func (EmployeeAsset) TableName() string {
	return "employee_assets"
}

// TableAliasName returns the table alias name of the EmployeeAsset table, used for querying.
func (EmployeeAsset) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the EmployeeAsset data in the database, used for querying.
func (m *EmployeeAsset) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "employees", "emp", []map[string]any{{"column1": "emp.id", "column2": "m.employee_id"}})
	m.AddRelation("left", "departments", "emp_dept", []map[string]any{{"column1": "emp_dept.id", "column2": "emp.department_id"}})
	m.AddRelation("left", "branches", "emp_brc", []map[string]any{{"column1": "emp_brc.id", "column2": "emp.branch_id"}})
	m.AddRelation("left", "attachments", "emp_att", []map[string]any{{"column1": "emp_att.id", "column2": "emp.attachment_id"}})

	m.AddRelation("left", "assets", "ass", []map[string]any{{"column1": "ass.id", "column2": "m.asset_id"}})
	m.AddRelation("left", "attachments", "ass_att", []map[string]any{{"column1": "ass_att.id", "column2": "ass.attachment_id"}})
	m.AddRelation("left", "categories", "ass_cat", []map[string]any{{"column1": "ass_cat.id", "column2": "ass.category_id"}})

	m.AddRelation("left", "attachments", "att", []map[string]any{{"column1": "att.id", "column2": "m.attachment_id"}})
	m.AddRelation("left", "conditions", "cond", []map[string]any{{"column1": "cond.id", "column2": "m.condition_id"}})

	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the EmployeeAsset data in the database, used for querying.
func (m *EmployeeAsset) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the EmployeeAsset data in the database, used for querying.
func (m *EmployeeAsset) GetSorts() []map[string]any {
	// m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the EmployeeAsset data in the database, used for querying.
func (m *EmployeeAsset) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the EmployeeAsset schema, used for querying.
func (m *EmployeeAsset) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the EmployeeAsset schema in the open api documentation.
func (EmployeeAsset) OpenAPISchemaName() string {
	return "EmployeeAsset"
}

// GetOpenAPISchema returns the Open API Schema of the EmployeeAsset in the open api documentation.
func (m *EmployeeAsset) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type EmployeeAssetList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the EmployeeAssetList schema in the open api documentation.
func (EmployeeAssetList) OpenAPISchemaName() string {
	return "EmployeeAssetList"
}

// GetOpenAPISchema returns the Open API Schema of the EmployeeAssetList in the open api documentation.
func (p *EmployeeAssetList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&EmployeeAsset{})
}

// ParamCreate is the expected parameters for create a new EmployeeAsset data.
type ParamCreate struct {
	UseCaseHandler
	AssignDate  app.NullDate `json:"assign_date"  db:"m.assign_date"  gorm:"column:assign_date"  validate:"required"`
	AssetID     app.NullUUID `json:"asset.id"     db:"m.asset_id"     gorm:"column:asset_id"     validate:"required"`
	EmployeeID  app.NullUUID `json:"employee.id"  db:"m.employee_id"  gorm:"column:employee_id"  validate:"required"`
	ConditionID app.NullUUID `json:"condition.id" db:"m.condition_id" gorm:"column:condition_id" validate:"required"`
}

// ParamUpdate is the expected parameters for update the EmployeeAsset data.
type ParamUpdate struct {
	UseCaseHandler
	AssignDate  app.NullDate `json:"assign_date"  db:"m.assign_date"  gorm:"column:assign_date"  validate:"required"`
	AssetID     app.NullUUID `json:"asset.id"     db:"m.asset_id"     gorm:"column:asset_id"     validate:"required"`
	EmployeeID  app.NullUUID `json:"employee.id"  db:"m.employee_id"  gorm:"column:employee_id"  validate:"required"`
	ConditionID app.NullUUID `json:"condition.id" db:"m.condition_id" gorm:"column:condition_id" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the EmployeeAsset data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the EmployeeAsset data.
type ParamDelete struct {
	UseCaseHandler
}
