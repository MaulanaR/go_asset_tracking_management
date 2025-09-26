package distributionassetsperdepartment

import (
	"net/http"
	"net/url"
	"time"

	"github.com/maulanar/go_asset_tracking_management/app"
)

// UseCase returns a UseCaseHandler for expected use case functional.
func UseCase(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	u := UseCaseHandler{
		Ctx:   &ctx,
		Query: url.Values{},
	}
	if len(query) > 0 {
		u.Query = query[0]
	}
	return u
}

// UseCaseHandler provides a convenient interface for DistributionAssetsPerDepartment use case, use UseCase to access UseCaseHandler.
type UseCaseHandler struct {
	DistributionAssetsPerDepartment

	// injectable dependencies
	Ctx   *app.Ctx   `json:"-" db:"-" gorm:"-"`
	Query url.Values `json:"-" db:"-" gorm:"-"`
}

// Async return UseCaseHandler with async process.
func (u UseCaseHandler) Async(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	ctx.IsAsync = true
	return UseCase(ctx, query...)
}

// Get returns the list of DistributionAssetsPerDepartment data.
func (u UseCaseHandler) Get() (res ViewData, err error) {

	// check permission
	err = u.Ctx.ValidatePermission("distribution_assets_per_departments.list")
	if err != nil {
		return res, err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// parsing filter
	deptFilter := u.Query.Get("department.id")
	catFilter := u.Query.Get("category.id")
	branchFilter := u.Query.Get("branch.id")

	// siapkan bagian WHERE tambahan
	where := ""
	args := []interface{}{}
	if deptFilter != "" {
		where += " AND d.id = ?"
		args = append(args, deptFilter)
	}
	if catFilter != "" {
		where += " AND c.id = ?"
		args = append(args, catFilter)
	}
	if branchFilter != "" {
		where += " AND b.id = ?"
		args = append(args, branchFilter)
	}

	rows := []DistributionAssetsPerDepartment{}
	query := `
WITH ea_latest AS (
  SELECT DISTINCT ON (ea.asset_id)
         ea.asset_id,
         ea.employee_id
  FROM employee_assets ea
  WHERE ea.deleted_at IS NULL
  ORDER BY ea.asset_id,
           COALESCE(ea.assign_date, ea.date, ea.created_at) DESC,
           ea.id DESC
)
SELECT
  d.id AS department_id,
  d.name AS department_name,
  c.id AS category_id,
  c.name AS category_name,
  b.id AS branch_id,
  b.name AS branch_name,
  COUNT(*) AS total_asset
FROM ea_latest x
JOIN assets a       ON a.id = x.asset_id        AND a.deleted_at IS NULL
JOIN employees e    ON e.id = x.employee_id     AND e.deleted_at IS NULL
LEFT JOIN departments d ON d.id = e.department_id AND d.deleted_at IS NULL
LEFT JOIN branches b    ON b.id = e.branch_id     AND b.deleted_at IS NULL
LEFT JOIN categories c  ON c.id = a.category_id   AND c.deleted_at IS NULL
WHERE 1=1` + where + `
GROUP BY d.id, d.name, c.id, c.name, b.id, b.name
ORDER BY d.name, c.name, b.name;
`
	err = tx.Raw(query, args...).Scan(&rows).Error
	if err != nil {
		return res, err
	}

	idx := map[string]int{}
	var groups []DeptGroup
	var grand int64

	for _, r := range rows {
		dept := r.DepartmentName.String // sesuaikan getter NullString Anda
		if dept == "" {
			dept = "-"
		}
		i, ok := idx[dept]
		if !ok {
			idx[dept] = len(groups)
			groups = append(groups, DeptGroup{Name: dept})
			i = idx[dept]
		}
		groups[i].Items = append(groups[i].Items, r)
		groups[i].Subtotal += r.TotalAsset.Int64 // sesuaikan getter NullInt64 Anda
		grand += r.TotalAsset.Int64
	}

	res = ViewData{
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		Groups:     groups,
		GrandTotal: grand,
	}

	return res, nil
}
