package distributionassetsperdepartment

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2/utils"
	"gorm.io/gorm"

	"github.com/maulanar/go_asset_tracking_management/app"
)

// prepareTest prepares the test.
func prepareTest(tb testing.TB) {
	app.Test()
	tx := app.Test().Tx
	app.DB().RegisterTable("main", DistributionAssetsPerDepartment{})
	app.DB().MigrateTable(tx, "main", app.Setting{})
	tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&DistributionAssetsPerDepartment{})

	app.Server().AddMiddleware(app.Test().NewCtx([]string{
		"distribution_assets_per_departments.detail",
		"distribution_assets_per_departments.list",
		"distribution_assets_per_departments.create",
		"distribution_assets_per_departments.edit",
		"distribution_assets_per_departments.delete",
	}))
	app.Server().AddRoute("/distribution_assets_per_departments", "GET", REST().Get, nil)
}

// getTestDistributionAssetsPerDepartmentID returns an available DistributionAssetsPerDepartment ID.
func getTestDistributionAssetsPerDepartmentID() string {
	return "todo"
}

// tests is test scenario.
var tests = []struct {
	description  string // description of the test case
	method       string // method to test
	path         string // route path to test
	token        string // token to test
	bodyRequest  string // body to test
	expectedCode int    // expected HTTP status code
	expectedBody string // expected body response
}{
	{
		description:  "Get empty list of DistributionAssetsPerDepartment",
		method:       "GET",
		path:         "/distribution_assets_per_departments",
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"count":0,"results":[]}`,
	},
	{
		description:  "Create DistributionAssetsPerDepartment with minimum payload",
		method:       "POST",
		path:         "/distribution_assets_per_departments",
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"name":"Kilogram"}`,
		expectedCode: http.StatusCreated,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Get DistributionAssetsPerDepartment by ID",
		method:       "GET",
		path:         "/distribution_assets_per_departments/" + getTestDistributionAssetsPerDepartmentID(),
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Update DistributionAssetsPerDepartment by ID",
		method:       "PUT",
		path:         "/distribution_assets_per_departments/" + getTestDistributionAssetsPerDepartmentID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Update DistributionAssetsPerDepartment by ID","name":"KG"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"KG"}`,
	},
	{
		description:  "Partially update DistributionAssetsPerDepartment by ID",
		method:       "PATCH",
		path:         "/distribution_assets_per_departments/" + getTestDistributionAssetsPerDepartmentID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Partially Update DistributionAssetsPerDepartment by ID","name":"Kilo Gram"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilo Gram"}`,
	},
	{
		description:  "Delete DistributionAssetsPerDepartment by ID",
		method:       "DELETE",
		path:         "/distribution_assets_per_departments/" + getTestDistributionAssetsPerDepartmentID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Delete DistributionAssetsPerDepartment by ID"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"code":200}`,
	},
}

// TestDistributionAssetsPerDepartmentREST tests the REST API of DistributionAssetsPerDepartment data with specified scenario.
func TestDistributionAssetsPerDepartmentREST(t *testing.T) {
	prepareTest(t)

	// Iterate through test single test cases
	for _, test := range tests {

		// Create a new http request with the route from the test case
		req := httptest.NewRequest(test.method, test.path, strings.NewReader(test.bodyRequest))
		req.Header.Add("Authorization", "Bearer "+test.token)
		req.Header.Add("Content-Type", "application/json")

		// Perform the request plain with the app, the second argument is a request latency (set to -1 for no latency)
		res, err := app.Server().Test(req)

		// Verify if the status code is as expected
		utils.AssertEqual(t, nil, err, "app.Server().Test(req)")
		utils.AssertEqual(t, test.expectedCode, res.StatusCode, test.description)

		// Verify if the body response is as expected
		body, err := io.ReadAll(res.Body)
		utils.AssertEqual(t, nil, err, "io.ReadAll(res.Body)")
		app.Test().AssertMatchJSONElement(t, []byte(test.expectedBody), body, test.description)
		res.Body.Close()
	}
}

// BenchmarkDistributionAssetsPerDepartmentREST tests the REST API of DistributionAssetsPerDepartment data with specified scenario.
func BenchmarkDistributionAssetsPerDepartmentREST(b *testing.B) {
	b.ReportAllocs()
	prepareTest(b)
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			req := httptest.NewRequest(test.method, test.path, strings.NewReader(test.bodyRequest))
			req.Header.Add("Authorization", "Bearer "+test.token)
			req.Header.Add("Content-Type", "application/json")
			app.Server().Test(req)
		}
	}
}
