package assetcondition

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
	app.DB().RegisterTable("main", AssetCondition{})
	app.DB().MigrateTable(tx, "main", app.Setting{})
	tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&AssetCondition{})

	app.Server().AddMiddleware(app.Test().NewCtx([]string{
		"asset_conditions.detail",
		"asset_conditions.list",
		"asset_conditions.create",
		"asset_conditions.edit",
		"asset_conditions.delete",
	}))
	app.Server().AddRoute("/asset_conditions", "GET", REST().Get, nil)
}

// getTestAssetConditionID returns an available AssetCondition ID.
func getTestAssetConditionID() string {
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
		description:  "Get empty list of AssetCondition",
		method:       "GET",
		path:         "/asset_conditions",
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"count":0,"results":[]}`,
	},
	{
		description:  "Create AssetCondition with minimum payload",
		method:       "POST",
		path:         "/asset_conditions",
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"name":"Kilogram"}`,
		expectedCode: http.StatusCreated,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Get AssetCondition by ID",
		method:       "GET",
		path:         "/asset_conditions/" + getTestAssetConditionID(),
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Update AssetCondition by ID",
		method:       "PUT",
		path:         "/asset_conditions/" + getTestAssetConditionID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Update AssetCondition by ID","name":"KG"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"KG"}`,
	},
	{
		description:  "Partially update AssetCondition by ID",
		method:       "PATCH",
		path:         "/asset_conditions/" + getTestAssetConditionID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Partially Update AssetCondition by ID","name":"Kilo Gram"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilo Gram"}`,
	},
	{
		description:  "Delete AssetCondition by ID",
		method:       "DELETE",
		path:         "/asset_conditions/" + getTestAssetConditionID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Delete AssetCondition by ID"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"code":200}`,
	},
}

// TestAssetConditionREST tests the REST API of AssetCondition data with specified scenario.
func TestAssetConditionREST(t *testing.T) {
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

// BenchmarkAssetConditionREST tests the REST API of AssetCondition data with specified scenario.
func BenchmarkAssetConditionREST(b *testing.B) {
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
