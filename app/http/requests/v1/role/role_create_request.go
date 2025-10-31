package role

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"karuhundeveloper.com/gostarterkit/app/helpers"
)

type RoleCreateRequest struct {
	Name string `form:"name" json:"name"`
}

func (r *RoleCreateRequest) Authorize(ctx http.Context) error {
	if !helpers.HasPermission(ctx, "create_role") {
		return helpers.NoPermissionError()
	}

	return nil
}

func (r *RoleCreateRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *RoleCreateRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name": "required|string|max_len:255|unique:roles,name",
	}
}

func (r *RoleCreateRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *RoleCreateRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *RoleCreateRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
