package user

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/maulanar/go_asset_tracking_management/app"
	"gorm.io/gorm"
)

// prepareTest prepares the test environment.
func prepareTest(tb testing.TB) {
	app.Test()
	tx := app.Test().Tx
	app.DB().RegisterTable("main", User{})
	app.DB().MigrateTable(tx, "main", app.Setting{})
	tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})

	app.Server().AddMiddleware(app.Test().NewCtx([]string{
		"users.register",
		"users.login",
	}))
	app.Server().AddRoute("/auth/register", "POST", REST().Register, nil)
	app.Server().AddRoute("/auth/login", "POST", REST().Login, nil)
}

// tests is test scenario.
var tests = []struct {
	description  string // description of the test case
	method       string // method to test
	path         string // route path to test
	bodyRequest  string // body to test
	expectedCode int    // expected HTTP status code
	expectedBody string // expected body response
}{
	{
		description:  "Register new user",
		method:       "POST",
		path:         "/auth/register",
		bodyRequest:  `{"email":"test@example.com","password":"secret"}`,
		expectedCode: http.StatusCreated,
		expectedBody: `{"message":"registered"}`,
	},
	{
		description:  "Login with correct credentials",
		method:       "POST",
		path:         "/auth/login",
		bodyRequest:  `{"email":"test@example.com","password":"secret"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"token":""}`, // token will be dynamic, so we just check element exists
	},
	{
		description:  "Login with wrong password",
		method:       "POST",
		path:         "/auth/login",
		bodyRequest:  `{"email":"test@example.com","password":"wrong"}`,
		expectedCode: http.StatusUnauthorized,
		expectedBody: `{"error":"invalid credentials"}`,
	},
}

// TestAuthREST tests the REST API of Auth module.
func TestAuthREST(t *testing.T) {
	prepareTest(t)

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.path, strings.NewReader(test.bodyRequest))
		req.Header.Add("Content-Type", "application/json")

		res, err := app.Server().Test(req)
		utils.AssertEqual(t, nil, err, "app.Server().Test(req)")
		utils.AssertEqual(t, test.expectedCode, res.StatusCode, test.description)

		body, err := io.ReadAll(res.Body)
		utils.AssertEqual(t, nil, err, "io.ReadAll(res.Body)")

		// special case: token is dynamic
		if strings.Contains(test.expectedBody, `"token":""`) {
			if !strings.Contains(string(body), `"token"`) {
				t.Errorf("%s: expected token in response, got %s", test.description, string(body))
			}
		} else {
			app.Test().AssertMatchJSONElement(t, []byte(test.expectedBody), body, test.description)
		}
		res.Body.Close()
	}
}

// BenchmarkAuthREST tests the performance of Auth REST API.
func BenchmarkAuthREST(b *testing.B) {
	b.ReportAllocs()
	prepareTest(b)
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			req := httptest.NewRequest(test.method, test.path, strings.NewReader(test.bodyRequest))
			req.Header.Add("Content-Type", "application/json")
			app.Server().Test(req)
		}
	}
}
