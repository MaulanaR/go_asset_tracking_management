package distributionassetsperdepartment

import "github.com/maulanar/go_asset_tracking_management/app"

// DistributionAssetsPerDepartment is the main model of DistributionAssetsPerDepartment data. It provides a convenient interface for app.ModelInterface
type DistributionAssetsPerDepartment struct {
	app.Model
	CategoryID     app.NullUUID   `json:"category.id"     db:"category_id"     gorm:"column:category_id"`
	CategoryName   app.NullString `json:"category.name"   db:"category_name"   gorm:"column:category_name"`
	DepartmentID   app.NullUUID   `json:"department.id"   db:"department_id"   gorm:"column:department_id"`
	DepartmentName app.NullString `json:"department.name" db:"department_name" gorm:"column:department_name"`
	BranchID       app.NullUUID   `json:"branch.id"       db:"branch_id"       gorm:"column:branch_id"`
	BranchName     app.NullString `json:"branch.name"     db:"branch_name"     gorm:"column:branch_name"`
	TotalAsset     app.NullInt64  `json:"total_asset"     db:"total_asset"     gorm:"column:total_asset"`
}

// view model agar template gampang
type DeptGroup struct {
	Name     string
	Items    []DistributionAssetsPerDepartment
	Subtotal int64
}

type ViewData struct {
	CreatedAt  string
	Groups     []DeptGroup
	GrandTotal int64
}

// EndPoint returns the DistributionAssetsPerDepartment end point, it used for cache key, etc.
func (DistributionAssetsPerDepartment) EndPoint() string {
	return "distribution_assets_per_departments"
}

// TableVersion returns the versions of the DistributionAssetsPerDepartment table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (DistributionAssetsPerDepartment) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the DistributionAssetsPerDepartment table in the database.
func (DistributionAssetsPerDepartment) TableName() string {
	return "distribution_assets_per_departments"
}

// TableAliasName returns the table alias name of the DistributionAssetsPerDepartment table, used for querying.
func (DistributionAssetsPerDepartment) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the DistributionAssetsPerDepartment data in the database, used for querying.
func (m *DistributionAssetsPerDepartment) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the DistributionAssetsPerDepartment data in the database, used for querying.
func (m *DistributionAssetsPerDepartment) GetFilters() []map[string]any {
	return m.Filters
}

// GetSorts returns the default sort of the DistributionAssetsPerDepartment data in the database, used for querying.
func (m *DistributionAssetsPerDepartment) GetSorts() []map[string]any {
	return m.Sorts
}

// GetFields returns list of the field of the DistributionAssetsPerDepartment data in the database, used for querying.
func (m *DistributionAssetsPerDepartment) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the DistributionAssetsPerDepartment schema, used for querying.
func (m *DistributionAssetsPerDepartment) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the DistributionAssetsPerDepartment schema in the open api documentation.
func (DistributionAssetsPerDepartment) OpenAPISchemaName() string {
	return "DistributionAssetsPerDepartment"
}

// GetOpenAPISchema returns the Open API Schema of the DistributionAssetsPerDepartment in the open api documentation.
func (m *DistributionAssetsPerDepartment) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type DistributionAssetsPerDepartmentList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the DistributionAssetsPerDepartmentList schema in the open api documentation.
func (DistributionAssetsPerDepartmentList) OpenAPISchemaName() string {
	return "DistributionAssetsPerDepartmentList"
}

// GetOpenAPISchema returns the Open API Schema of the DistributionAssetsPerDepartmentList in the open api documentation.
func (p *DistributionAssetsPerDepartmentList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&DistributionAssetsPerDepartment{})
}

// ParamCreate is the expected parameters for create a new DistributionAssetsPerDepartment data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the DistributionAssetsPerDepartment data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the DistributionAssetsPerDepartment data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the DistributionAssetsPerDepartment data.
type ParamDelete struct {
	UseCaseHandler
}
