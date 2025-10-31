package role

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type RoleUpdateRequest struct {
	Name string `form:"name" json:"name"`
}

func (r *RoleUpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *RoleUpdateRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *RoleUpdateRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name": "required|string|max_len:255|unique:roles,name," + ctx.Request().Route("id"),
	}
}

func (r *RoleUpdateRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *RoleUpdateRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *RoleUpdateRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
