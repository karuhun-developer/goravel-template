package role

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"karuhundeveloper.com/gostarterkit/app/helpers"
)

type PermissionCreateRequest struct {
	Name string `form:"name" json:"name"`
}

func (r *PermissionCreateRequest) Authorize(ctx http.Context) error {
	if !helpers.HasPermission(ctx, "create_permission") {
		return helpers.NoPermissionError()
	}

	return nil
}

func (r *PermissionCreateRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PermissionCreateRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name": "required|string|max_len:255|unique:permissions,name",
	}
}

func (r *PermissionCreateRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PermissionCreateRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PermissionCreateRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
