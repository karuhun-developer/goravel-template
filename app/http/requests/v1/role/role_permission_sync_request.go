package role

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"karuhundeveloper.com/gostarterkit/app/helpers"
)

type RolePermssionSyncRequest struct {
	Permissions string `form:"permissions" json:"permissions"`
}

func (r *RolePermssionSyncRequest) Authorize(ctx http.Context) error {
	if !helpers.HasPermission(ctx, "update_role") {
		return helpers.NoPermissionError()
	}

	return nil
}

func (r *RolePermssionSyncRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *RolePermssionSyncRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"permissions": "required|array",
	}
}

func (r *RolePermssionSyncRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *RolePermssionSyncRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *RolePermssionSyncRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
