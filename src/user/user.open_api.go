package user

import "github.com/maulanar/go_asset_tracking_management/app"

func OpenAPI() *OpenAPIOperation {
	return &OpenAPIOperation{}
}

type OpenAPIOperation struct {
	app.OpenAPIOperation
}

func (o *OpenAPIOperation) Base() {
	o.Tags = []string{"User"}
	o.HeaderParams = []map[string]any{{"$ref": "#/components/parameters/headerParam.Accept-Language"}}
	o.Responses = map[string]map[string]any{
		"200": {
			"description": "Success",
			"content":     map[string]any{"application/json": &User{}}, // will auto create schema $ref: '#/components/schemas/Department' if not exists
		},
		"400": app.OpenAPIError().BadRequest(),
		"401": app.OpenAPIError().Unauthorized(),
		"403": app.OpenAPIError().Forbidden(),
	}
	o.Securities = []map[string][]string{}
}

// GetAll dokumentasi OpenAPI untuk mengambil semua data user
func (o *OpenAPIOperation) Get() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o
	}

	o.Base()
	o.Summary = "Get Users"
	o.Description = "Use this method to get list of Users"
	o.QueryParams = []map[string]any{{"$ref": "#/components/parameters/queryParam.Any"}}
	o.Responses = map[string]map[string]any{
		"200": {
			"description": "Success",
			"content":     map[string]any{"application/json": &UserList{}}, // will auto create schema $ref: '#/components/schemas/Department.List' if not exists
		},
		"400": app.OpenAPIError().BadRequest(),
		"401": app.OpenAPIError().Unauthorized(),
		"403": app.OpenAPIError().Forbidden(),
	}
	return o
}

// DeleteByID dokumentasi OpenAPI untuk menghapus user berdasarkan ID
func (o *OpenAPIOperation) DeleteByID() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o
	}

	o.Base()
	o.Tags = []string{"User"}
	o.Summary = "Delete User by ID"
	o.Description = "Delete a specific user record by its ID."
	o.PathParams = []map[string]any{
		{
			"name":     "id",
			"in":       "path",
			"required": true,
			"schema": map[string]string{
				"type": "string",
			},
			"description": "UUID of the user to delete",
		},
	}
	o.Securities = []map[string][]string{
		{"BearerAuth": {}},
	}
	o.Responses = map[string]map[string]any{
		"200": {
			"description": "Success",
			"content": map[string]any{
				"application/json": map[string]string{
					"message": "user deleted successfully",
				},
			},
		},
		"400": app.OpenAPIError().BadRequest(),
		"401": app.OpenAPIError().Unauthorized(),
		"403": app.OpenAPIError().Forbidden(),
	}
	return o
}

func (o *OpenAPIOperation) Register() *OpenAPIOperation {
	o.Base()
	o.Summary = "Register User"
	o.Description = "Register a new user"
	o.Body = map[string]any{"application/json": &ParamRegister{}}
	return o
}

func (o *OpenAPIOperation) Login() *OpenAPIOperation {
	o.Base()
	o.Summary = "Login User"
	o.Description = "Login with email and password"
	o.Body = map[string]any{"application/json": &ParamLogin{}}
	return o
}

func (o *OpenAPIOperation) Profile() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o
	}

	o.Base()
	o.Summary = "Get Profile"
	o.Description = "Get current logged in user profile"
	o.Responses = map[string]map[string]any{
		"200": {
			"description": "Success",
			"content":     map[string]any{"application/json": &User{}},
		},
		"401": app.OpenAPIError().Unauthorized(),
	}
	return o
}
