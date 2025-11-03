package role

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"karuhundeveloper.com/gostarterkit/app/helpers"
)

type PermissionUpdateRequest struct {
	Name string `form:"name" json:"name"`
}

func (r *PermissionUpdateRequest) Authorize(ctx http.Context) error {
	if !helpers.HasPermission(ctx, "update_permission") {
		return helpers.NoPermissionError()
	}

	return nil
}

func (r *PermissionUpdateRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PermissionUpdateRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name": "required|string|max_len:255|unique:permissions,name," + ctx.Request().Route("id"),
	}
}

func (r *PermissionUpdateRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PermissionUpdateRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PermissionUpdateRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
