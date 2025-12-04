package role

import (
	"github.com/goravel/framework/contracts/http"
	"karuhundeveloper.com/gostarterkit/app/helpers"
	request "karuhundeveloper.com/gostarterkit/app/http/requests/v1/role"
	"karuhundeveloper.com/gostarterkit/app/services/v1/role"
)

type RolePermissionController struct {
	// Dependent services
	rolePermissionService *role.RolePermissionService
}

func NewRolePermissionController(RolePermissionService *role.RolePermissionService) *RolePermissionController {
	return &RolePermissionController{
		rolePermissionService: RolePermissionService,
	}
}

func (r *RolePermissionController) Show(ctx http.Context) http.Response {
	if !helpers.HasPermission(ctx, "view_role") {
		return helpers.ErrorResponse(ctx, http.StatusForbidden, "Validation Error", helpers.NoPermissionError().Error())
	}

	// Show role
	role, err := r.rolePermissionService.Show(ctx)

	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusNotFound, "Role not found", err.Error())
	}

	return helpers.SuccessResponse(ctx, http.StatusOK, "Role Retrieved Successfully", helpers.ModelToMap(role))
}

func (r *RolePermissionController) Sync(ctx http.Context) http.Response {
	var rolePermssionSyncRequest request.RolePermssionSyncRequest

	// Validate request data
	validationErrors, err := ctx.Request().ValidateRequest(&rolePermssionSyncRequest)
	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Validation Error", err.Error())
	}

	if validationErrors != nil {
		return helpers.ErrorValidationResponse(ctx, http.StatusUnprocessableEntity, "Validation Error", validationErrors.All())
	}

	// Create role
	role, err := r.rolePermissionService.Sync(ctx, rolePermssionSyncRequest)

	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusBadRequest, "Role Creation Failed", err.Error())
	}

	return helpers.SuccessResponse(ctx, http.StatusCreated, "Role Permission Synced Successfully", helpers.ModelToMap(role))
}