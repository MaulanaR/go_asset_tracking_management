package src

import (
	"github.com/maulanar/go_asset_tracking_management/app"
	"github.com/maulanar/go_asset_tracking_management/src/asset"
	"github.com/maulanar/go_asset_tracking_management/src/attachment"
	"github.com/maulanar/go_asset_tracking_management/src/branch"
	"github.com/maulanar/go_asset_tracking_management/src/category"
	"github.com/maulanar/go_asset_tracking_management/src/condition"
	"github.com/maulanar/go_asset_tracking_management/src/department"
	"github.com/maulanar/go_asset_tracking_management/src/employee"
	"github.com/maulanar/go_asset_tracking_management/src/employeeasset"
	"github.com/maulanar/go_asset_tracking_management/src/jobposition"
	"github.com/maulanar/go_asset_tracking_management/src/reports/assetcondition"
	"github.com/maulanar/go_asset_tracking_management/src/reports/distributionassetsperdepartment"
	// import : DONT REMOVE THIS COMMENT
)

func Router() *routerUtil {
	if router == nil {
		router = &routerUtil{}
		router.Configure()
		router.isConfigured = true
	}
	return router
}

var router *routerUtil

type routerUtil struct {
	isConfigured bool
}

func (r *routerUtil) Configure() {
	app.Server().AddRoute("/api/version", "GET", app.VersionHandler, nil)
	app.Server().AddRoute("/api/v1/auth/me", "GET", app.VersionHandler, nil)

	app.Server().AddRoute("/api/v1/departments", "POST", department.REST().Create, department.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/departments", "GET", department.REST().Get, department.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/departments/{id}", "GET", department.REST().GetByID, department.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/departments/{id}", "PUT", department.REST().UpdateByID, department.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/departments/{id}", "PATCH", department.REST().PartiallyUpdateByID, department.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/departments/{id}", "DELETE", department.REST().DeleteByID, department.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/conditions", "POST", condition.REST().Create, condition.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/conditions", "GET", condition.REST().Get, condition.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/conditions/{id}", "GET", condition.REST().GetByID, condition.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/conditions/{id}", "PUT", condition.REST().UpdateByID, condition.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/conditions/{id}", "PATCH", condition.REST().PartiallyUpdateByID, condition.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/conditions/{id}", "DELETE", condition.REST().DeleteByID, condition.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/categories", "POST", category.REST().Create, category.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/categories", "GET", category.REST().Get, category.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/categories/{id}", "GET", category.REST().GetByID, category.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/categories/{id}", "PUT", category.REST().UpdateByID, category.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/categories/{id}", "PATCH", category.REST().PartiallyUpdateByID, category.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/categories/{id}", "DELETE", category.REST().DeleteByID, category.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/branches", "POST", branch.REST().Create, branch.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/branches", "GET", branch.REST().Get, branch.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/branches/{id}", "GET", branch.REST().GetByID, branch.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/branches/{id}", "PUT", branch.REST().UpdateByID, branch.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/branches/{id}", "PATCH", branch.REST().PartiallyUpdateByID, branch.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/branches/{id}", "DELETE", branch.REST().DeleteByID, branch.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/employees", "POST", employee.REST().Create, employee.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/employees", "GET", employee.REST().Get, employee.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/employees/{id}", "GET", employee.REST().GetByID, employee.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/employees/{id}", "PUT", employee.REST().UpdateByID, employee.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/employees/{id}", "PATCH", employee.REST().PartiallyUpdateByID, employee.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/employees/{id}", "DELETE", employee.REST().DeleteByID, employee.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/assets", "POST", asset.REST().Create, asset.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/assets", "GET", asset.REST().Get, asset.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/assets/{id}", "GET", asset.REST().GetByID, asset.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/assets/{id}/depreciations", "GET", asset.REST().GetDepreciationByID, asset.OpenAPI().GetDepreciationByID())
	app.Server().AddRoute("/api/v1/assets/{id}", "PUT", asset.REST().UpdateByID, asset.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/assets/{id}", "PATCH", asset.REST().PartiallyUpdateByID, asset.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/assets/{id}", "DELETE", asset.REST().DeleteByID, asset.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/attachments", "POST", attachment.REST().Create, attachment.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/attachments", "GET", attachment.REST().Get, attachment.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/attachments/{id}", "GET", attachment.REST().GetByID, attachment.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/attachments/{id}", "DELETE", attachment.REST().DeleteByID, attachment.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/employee_assets", "POST", employeeasset.REST().Create, employeeasset.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/employee_assets", "GET", employeeasset.REST().Get, employeeasset.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/employee_assets/{id}", "GET", employeeasset.REST().GetByID, employeeasset.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/employee_assets/{id}", "PUT", employeeasset.REST().UpdateByID, employeeasset.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/employee_assets/{id}", "PATCH", employeeasset.REST().PartiallyUpdateByID, employeeasset.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/employee_assets/{id}", "DELETE", employeeasset.REST().DeleteByID, employeeasset.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/job_positions", "POST", jobposition.REST().Create, jobposition.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/job_positions", "GET", jobposition.REST().Get, jobposition.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/job_positions/{id}", "GET", jobposition.REST().GetByID, jobposition.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/job_positions/{id}", "PUT", jobposition.REST().UpdateByID, jobposition.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/job_positions/{id}", "PATCH", jobposition.REST().PartiallyUpdateByID, jobposition.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/job_positions/{id}", "DELETE", jobposition.REST().DeleteByID, jobposition.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/reports/distribution_assets_per_departments", "GET", distributionassetsperdepartment.REST().Get, distributionassetsperdepartment.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/reports/asset_conditions", "GET", assetcondition.REST().Get, assetcondition.OpenAPI().Get())

	// AddRoute : DONT REMOVE THIS COMMENT
}
