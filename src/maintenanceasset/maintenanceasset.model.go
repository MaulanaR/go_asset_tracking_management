package maintenanceasset

import "github.com/maulanar/go_asset_tracking_management/app"

// MaintenanceAsset is the main model of MaintenanceAsset data. It provides a convenient interface for app.ModelInterface
type MaintenanceAsset struct {
	app.Model
	ID          app.NullUUID    `json:"id"                                db:"m.id"                              gorm:"column:id;primaryKey"`
	Code        app.NullString  `json:"code"                              db:"m.code"                            gorm:"column:code"`
	Date        app.NullDate    `json:"date"                              db:"m.date"                            gorm:"column:date"`
	Description app.NullText    `json:"description"                       db:"m.description"                     gorm:"column:description"`
	Amount      app.NullFloat64 `json:"amount"                            db:"m.amount"                          gorm:"column:amount"`

	AssetID                         app.NullUUID    `json:"asset.id"                          db:"m.asset_id"                      gorm:"column:asset_id"`
	AssetCode                       app.NullString  `json:"asset.code"                        db:"ass.code"                          gorm:"-"`
	AssetName                       app.NullString  `json:"asset.name"                        db:"ass.name"                          gorm:"-"`
	AssetInputDate                  app.NullDate    `json:"asset.input_date"                  db:"ass.input_date"                    gorm:"-"`
	AssetPrice                      app.NullFloat64 `json:"asset.price"                       db:"ass.price"                         gorm:"-"`
	AssetAttachmentID               app.NullUUID    `json:"asset.attachment.id"               db:"ass.attachment_id"                 gorm:"-"`
	AssetAttachmentName             app.NullText    `json:"asset.attachment.name"             db:"assatt.name"                          gorm:"-"`
	AssetAttachmentPath             app.NullText    `json:"asset.attachment.path"             db:"assatt.path"                          gorm:"-"`
	AssetAttachmentURL              app.NullText    `json:"asset.attachment.url"              db:"assatt.url"                           gorm:"-"`
	AssetCategoryID                 app.NullUUID    `json:"asset.category.id"                 db:"ass.category_id"                   gorm:"-"`
	AssetCategoryCode               app.NullString  `json:"asset.category.code"               db:"asscat.code"                          gorm:"-"`
	AssetCategoryName               app.NullString  `json:"asset.category.name"               db:"asscat.name"                          gorm:"-"`
	AssetCategoryEconomicAges       app.NullInt64   `json:"asset.category.economic_age"       db:"asscat.economic_age"                  gorm:"-"`
	AssetCategoryDescription        app.NullText    `json:"asset.category.description"        db:"asscat.description"                   gorm:"-"`
	AssetAssignDate                 app.NullDate    `json:"asset.assign_date"                 db:"emp_ass.assign_date"               gorm:"-"`
	AssetConditionID                app.NullUUID    `json:"asset.condition.id"                db:"emp_ass.condition_id"              gorm:"-"`
	AssetConditionCode              app.NullString  `json:"asset.condition.code"              db:"emp_ass_cond.code"                 gorm:"-"`
	AssetConditionName              app.NullString  `json:"asset.condition.name"              db:"emp_ass_cond.name"                 gorm:"-"`
	AssetConditionDescription       app.NullText    `json:"asset.condition.description"       db:"emp_ass_cond.description"          gorm:"-"`
	AssetDepartmentID               app.NullUUID    `json:"asset.department.id"               db:"emp_ass_dpt.id"                    gorm:"-"`
	AssetDepartmentCode             app.NullString  `json:"asset.department.code"             db:"emp_ass_dpt.code"                  gorm:"-"`
	AssetDepartmentName             app.NullString  `json:"asset.department.name"             db:"emp_ass_dpt.name"                  gorm:"-"`
	AssetDepartmentDescription      app.NullText    `json:"asset.department.description"      db:"emp_ass_dpt.description"           gorm:"-"`
	AssetEmployeeID                 app.NullUUID    `json:"asset.employee.id"                 db:"emp_ass_emp.id"                            gorm:"-"`
	AssetEmployeeCode               app.NullString  `json:"asset.employee.code"               db:"emp_ass_emp.code"                          gorm:"-"`
	AssetEmployeeName               app.NullString  `json:"asset.employee.name"               db:"emp_ass_emp.name"                          gorm:"-"`
	AssetEmployeeAddress            app.NullText    `json:"asset.employee.address"            db:"emp_ass_emp.address"                       gorm:"-"`
	AssetEmployeePhone              app.NullString  `json:"asset.employee.phone"              db:"emp_ass_emp.phone"                         gorm:"-"`
	AssetEmployeeEmail              app.NullString  `json:"asset.employee.email"              db:"emp_ass_emp.email"                         gorm:"-"`
	AssetEmployeeIsActive           app.NullBool    `json:"asset.employee.is_active"          db:"emp_ass_emp.is_active"                     gorm:"-"`
	AssetBranchID                   app.NullUUID    `json:"asset.branch.id"                   db:"emp_ass_brc.id"                    gorm:"-"`
	AssetBranchCode                 app.NullString  `json:"asset.branch.code"                 db:"emp_ass_brc.code"                  gorm:"-"`
	AssetBranchName                 app.NullString  `json:"asset.branch.name"                 db:"emp_ass_brc.name"                  gorm:"-"`
	AssetBranchAddress              app.NullText    `json:"asset.branch.address"              db:"emp_ass_brc.address"               gorm:"-"`
	AssetStatus                     app.NullString  `json:"asset.status"                      db:"ass.status"                        gorm:"-"`
	AssetDepreciationAmount         app.NullFloat64 `json:"asset.depreciation.amount"         db:"ass.depreciation_amount"           gorm:"-"`
	AssetDepreciationAmountPerMonth app.NullFloat64 `json:"asset.depreciation.per_month"      db:"ass.depreciation_amount_per_month" gorm:"-"`
	AssetSalvageAmount              app.NullFloat64 `json:"asset.salvage.amount"              db:"ass.salvage_amount"                gorm:"-"`
	AssetCurrentValue               app.NullFloat64 `json:"asset.current.amount"              db:"ass.current_amount"                gorm:"-"`

	MaintenanceTypeID          app.NullUUID   `json:"maintenance_type.id"               db:"m.maintenance_type_id"             gorm:"column:maintenance_type_id"`
	MaintenanceTypeCode        app.NullString `json:"maintenance_type.code"             db:"mt.code"                           gorm:"-"`
	MaintenanceTypeName        app.NullString `json:"maintenance_type.name"             db:"mt.name"                           gorm:"-"`
	MaintenanceTypeDescription app.NullText   `json:"maintenance_type.description"      db:"mt.description"                    gorm:"-"`
	MaintenanceTypeIsActive    app.NullBool   `json:"maintenance_type.is_active"        db:"mt.is_active"                      gorm:"-"`

	EmployeeId                     app.NullUUID   `json:"employee.id"                       db:"m.employee_id"                     gorm:"column:employee_id"`
	EmployeeCode                   app.NullString `json:"employee.code"                     db:"emp.code"                          gorm:"-"`
	EmployeeName                   app.NullString `json:"employee.name"                     db:"emp.name"                          gorm:"-"`
	EmployeeDepartmentID           app.NullUUID   `json:"employee.department.id"            db:"emp.department_id"                 gorm:"-"`
	EmployeeDepartmentCode         app.NullString `json:"employee.department.code"          db:"empdpt.code"                       gorm:"-"`
	EmployeeDepartmentName         app.NullString `json:"employee.department.name"          db:"empdpt.name"                       gorm:"-"`
	EmployeeDepartmentDescription  app.NullText   `json:"employee.department.description"   db:"empdpt.description"                gorm:"-"`
	EmployeeBranchID               app.NullUUID   `json:"employee.branch.id"                db:"emp.branch_id"                     gorm:"-"`
	EmployeeBranchCode             app.NullString `json:"employee.branch.code"              db:"empbrc.code"                       gorm:"-"`
	EmployeeBranchName             app.NullString `json:"employee.branch.name"              db:"empbrc.name"                       gorm:"-"`
	EmployeeBranchAddress          app.NullText   `json:"employee.branch.address"           db:"empbrc.address"                    gorm:"-"`
	EmployeeAddress                app.NullText   `json:"employee.address"                  db:"emp.address"                       gorm:"-"`
	EmployeePhone                  app.NullString `json:"employee.phone"                    db:"emp.phone"                         gorm:"-"`
	EmployeeAttachmentID           app.NullUUID   `json:"employee.attachment.id"            db:"emp.attachment_id"                 gorm:"-"`
	EmployeeAttachmentName         app.NullText   `json:"employee.attachment.name"          db:"empatt.name"                       gorm:"-"`
	EmployeeAttachmentPath         app.NullText   `json:"employee.attachment.path"          db:"empatt.path"                       gorm:"-"`
	EmployeeAttachmentURL          app.NullText   `json:"employee.attachment.url"           db:"empatt.url"                        gorm:"-"`
	EmployeeEmail                  app.NullString `json:"employee.email"                    db:"emp.email"                         gorm:"-"`
	EmployeeType                   app.NullString `json:"employee.type"                     db:"emp.type"                          gorm:"-"`
	EmployeeJobPositionID          app.NullUUID   `json:"employee.job_position.id"          db:"emp.job_position_id"               gorm:"-"`
	EmployeeJobPositionCode        app.NullString `json:"employee.job_position.code"        db:"empjp.code"                        gorm:"-"`
	EmployeeJobPositionName        app.NullString `json:"employee.job_position.name"        db:"empjp.name"                        gorm:"-"`
	EmployeeJobPositionDescription app.NullText   `json:"employee.job_position.description" db:"empjp.description"                 gorm:"-"`
	EmployeeIsActive               app.NullBool   `json:"employee.is_active"                db:"emp.is_active"                     gorm:"-"`

	AttachmentId   app.NullUUID `json:"attachment.id"                     db:"m.attachment_id"                   gorm:"column:attachment_id"`
	AttachmentName app.NullText `json:"attachment.name"                   db:"att.name"                          gorm:"-"`
	AttachmentPath app.NullText `json:"attachment.path"                   db:"att.path"                          gorm:"-"`
	AttachmentURL  app.NullText `json:"attachment.url"                    db:"att.url"                           gorm:"-"`

	CreatedAt app.NullDateTime `json:"created_at"                        db:"m.created_at"                      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at"                        db:"m.updated_at"                      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at"                        db:"m.deleted_at,hide"                 gorm:"column:deleted_at"`
}

// EndPoint returns the MaintenanceAsset end point, it used for cache key, etc.
func (MaintenanceAsset) EndPoint() string {
	return "maintenance_assets"
}

// TableVersion returns the versions of the MaintenanceAsset table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (MaintenanceAsset) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the MaintenanceAsset table in the database.
func (MaintenanceAsset) TableName() string {
	return "maintenance_assets"
}

// TableAliasName returns the table alias name of the MaintenanceAsset table, used for querying.
func (MaintenanceAsset) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the MaintenanceAsset data in the database, used for querying.
func (m *MaintenanceAsset) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "assets", "ass", []map[string]any{{"column1": "ass.id", "column2": "m.asset_id"}})
	m.AddRelation("left", "attachments", "assatt", []map[string]any{{"column1": "assatt.id", "column2": "ass.attachment_id"}})
	m.AddRelation("left", "categories", "asscat", []map[string]any{{"column1": "asscat.id", "column2": "ass.category_id"}})

	// search to employee_assets
	m.AddRelation("left", `(
  SELECT DISTINCT ON (ea.asset_id)
         ea.assign_date,
         ea.asset_id,
         ea.condition_id,
         ea.employee_id
  FROM employee_assets ea
  WHERE ea.deleted_at IS NULL
  ORDER BY ea.asset_id, ea.assign_date DESC, ea.id DESC
)`, "emp_ass", []map[string]any{{"column1": "emp_ass.asset_id", "column2": "m.id"}})

	m.AddRelation("left", "conditions", "emp_ass_cond", []map[string]any{{"column1": "emp_ass_cond.id", "column2": "emp_ass.condition_id"}})
	m.AddRelation("left", "employees", "emp_ass_emp", []map[string]any{{"column1": "emp_ass_emp.id", "column2": "emp_ass.employee_id"}})
	m.AddRelation("left", "departments", "emp_ass_dpt", []map[string]any{{"column1": "emp_ass_dpt.id", "column2": "emp_ass_emp.department_id"}})
	m.AddRelation("left", "branches", "emp_ass_brc", []map[string]any{{"column1": "emp_ass_brc.id", "column2": "emp_ass_emp.branch_id"}})

	m.AddRelation("left", "maintenance_types", "mt", []map[string]any{{"column1": "mt.id", "column2": "m.maintenance_type_id"}})
	m.AddRelation("left", "employees", "emp", []map[string]any{{"column1": "emp.id", "column2": "m.employee_id"}})
	m.AddRelation("left", "departments", "empdpt", []map[string]any{{"column1": "empdpt.id", "column2": "emp.department_id"}})
	m.AddRelation("left", "branches", "empbrc", []map[string]any{{"column1": "empbrc.id", "column2": "emp.branch_id"}})
	m.AddRelation("left", "attachments", "empatt", []map[string]any{{"column1": "empatt.id", "column2": "emp.attachment_id"}})
	m.AddRelation("left", "job_positions", "empjp", []map[string]any{{"column1": "empjp.id", "column2": "emp.job_position_id"}})

	m.AddRelation("left", "attachments", "att", []map[string]any{{"column1": "att.id", "column2": "m.attachment_id"}})
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the MaintenanceAsset data in the database, used for querying.
func (m *MaintenanceAsset) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the MaintenanceAsset data in the database, used for querying.
func (m *MaintenanceAsset) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the MaintenanceAsset data in the database, used for querying.
func (m *MaintenanceAsset) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the MaintenanceAsset schema, used for querying.
func (m *MaintenanceAsset) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the MaintenanceAsset schema in the open api documentation.
func (MaintenanceAsset) OpenAPISchemaName() string {
	return "MaintenanceAsset"
}

// GetOpenAPISchema returns the Open API Schema of the MaintenanceAsset in the open api documentation.
func (m *MaintenanceAsset) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type MaintenanceAssetList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the MaintenanceAssetList schema in the open api documentation.
func (MaintenanceAssetList) OpenAPISchemaName() string {
	return "MaintenanceAssetList"
}

// GetOpenAPISchema returns the Open API Schema of the MaintenanceAssetList in the open api documentation.
func (p *MaintenanceAssetList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&MaintenanceAsset{})
}

// ParamCreate is the expected parameters for create a new MaintenanceAsset data.
type ParamCreate struct {
	UseCaseHandler
	Date       app.NullDate    `json:"date"                              db:"m.date"                            gorm:"column:date" validate:"required"`
	Amount     app.NullFloat64 `json:"amount"                            db:"m.amount"                          gorm:"column:amount" validate:"required"`
	AssetID    app.NullUUID    `json:"asset.id"                          db:"m.asset_id"                      gorm:"column:asset_id" validate:"required"`
	EmployeeId app.NullUUID    `json:"employee.id"                       db:"m.employee_id"                     gorm:"column:employee_id" validate:"required"`
}

// ParamUpdate is the expected parameters for update the MaintenanceAsset data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the MaintenanceAsset data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the MaintenanceAsset data.
type ParamDelete struct {
	UseCaseHandler
}
