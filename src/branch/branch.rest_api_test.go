package branch

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
	app.DB().RegisterTable("main", Branch{})
	app.DB().MigrateTable(tx, "main", app.Setting{})
	tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Branch{})

	app.Server().AddMiddleware(app.Test().NewCtx([]string{
		"branches.detail",
		"branches.list",
		"branches.create",
		"branches.edit",
		"branches.delete",
	}))
	app.Server().AddRoute("/branches", "POST", REST().Create, nil)
	app.Server().AddRoute("/branches", "GET", REST().Get, nil)
	app.Server().AddRoute("/branches/:id", "GET", REST().GetByID, nil)
	app.Server().AddRoute("/branches/:id", "PUT", REST().UpdateByID, nil)
	app.Server().AddRoute("/branches/:id", "PATCH", REST().PartiallyUpdateByID, nil)
	app.Server().AddRoute("/branches/:id", "DELETE", REST().DeleteByID, nil)
}

// getTestBranchID returns an available Branch ID.
func getTestBranchID() string {
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
		description:  "Get empty list of Branch",
		method:       "GET",
		path:         "/branches",
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"count":0,"results":[]}`,
	},
	{
		description:  "Create Branch with minimum payload",
		method:       "POST",
		path:         "/branches",
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"name":"Kilogram"}`,
		expectedCode: http.StatusCreated,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Get Branch by ID",
		method:       "GET",
		path:         "/branches/" + getTestBranchID(),
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Update Branch by ID",
		method:       "PUT",
		path:         "/branches/" + getTestBranchID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Update Branch by ID","name":"KG"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"KG"}`,
	},
	{
		description:  "Partially update Branch by ID",
		method:       "PATCH",
		path:         "/branches/" + getTestBranchID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Partially Update Branch by ID","name":"Kilo Gram"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilo Gram"}`,
	},
	{
		description:  "Delete Branch by ID",
		method:       "DELETE",
		path:         "/branches/" + getTestBranchID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Delete Branch by ID"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"code":200}`,
	},
}

// TestBranchREST tests the REST API of Branch data with specified scenario.
func TestBranchREST(t *testing.T) {
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

// BenchmarkBranchREST tests the REST API of Branch data with specified scenario.
func BenchmarkBranchREST(b *testing.B) {
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
