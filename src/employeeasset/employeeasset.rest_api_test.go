package employeeasset

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
	app.DB().RegisterTable("main", EmployeeAsset{})
	app.DB().MigrateTable(tx, "main", app.Setting{})
	tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&EmployeeAsset{})

	app.Server().AddMiddleware(app.Test().NewCtx([]string{
		"employee_assets.detail",
		"employee_assets.list",
		"employee_assets.create",
		"employee_assets.edit",
		"employee_assets.delete",
	}))
	app.Server().AddRoute("/employee_assets", "POST", REST().Create, nil)
	app.Server().AddRoute("/employee_assets", "GET", REST().Get, nil)
	app.Server().AddRoute("/employee_assets/:id", "GET", REST().GetByID, nil)
	app.Server().AddRoute("/employee_assets/:id", "PUT", REST().UpdateByID, nil)
	app.Server().AddRoute("/employee_assets/:id", "PATCH", REST().PartiallyUpdateByID, nil)
	app.Server().AddRoute("/employee_assets/:id", "DELETE", REST().DeleteByID, nil)
}

// getTestEmployeeAssetID returns an available EmployeeAsset ID.
func getTestEmployeeAssetID() string {
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
		description:  "Get empty list of EmployeeAsset",
		method:       "GET",
		path:         "/employee_assets",
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"count":0,"results":[]}`,
	},
	{
		description:  "Create EmployeeAsset with minimum payload",
		method:       "POST",
		path:         "/employee_assets",
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"name":"Kilogram"}`,
		expectedCode: http.StatusCreated,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Get EmployeeAsset by ID",
		method:       "GET",
		path:         "/employee_assets/" + getTestEmployeeAssetID(),
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Update EmployeeAsset by ID",
		method:       "PUT",
		path:         "/employee_assets/" + getTestEmployeeAssetID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Update EmployeeAsset by ID","name":"KG"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"KG"}`,
	},
	{
		description:  "Partially update EmployeeAsset by ID",
		method:       "PATCH",
		path:         "/employee_assets/" + getTestEmployeeAssetID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Partially Update EmployeeAsset by ID","name":"Kilo Gram"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilo Gram"}`,
	},
	{
		description:  "Delete EmployeeAsset by ID",
		method:       "DELETE",
		path:         "/employee_assets/" + getTestEmployeeAssetID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Delete EmployeeAsset by ID"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"code":200}`,
	},
}

// TestEmployeeAssetREST tests the REST API of EmployeeAsset data with specified scenario.
func TestEmployeeAssetREST(t *testing.T) {
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

// BenchmarkEmployeeAssetREST tests the REST API of EmployeeAsset data with specified scenario.
func BenchmarkEmployeeAssetREST(b *testing.B) {
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
