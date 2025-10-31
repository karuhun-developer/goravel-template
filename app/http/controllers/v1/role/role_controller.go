package role

import (
	"github.com/goravel/framework/contracts/http"
	"karuhundeveloper.com/gostarterkit/app/helpers"
	request "karuhundeveloper.com/gostarterkit/app/http/requests/v1/role"
	"karuhundeveloper.com/gostarterkit/app/services/v1/role"
)

type RoleController struct {
	// Dependent services
	roleService *role.RoleService
}

func NewRoleController(RoleService *role.RoleService) *RoleController {
	return &RoleController{
		roleService: RoleService,
	}
}

func (r *RoleController) Index(ctx http.Context) http.Response {
	if !helpers.HasPermission(ctx, "view_all_role") {
		return helpers.ErrorResponse(ctx, http.StatusForbidden, "Validation Error", helpers.NoPermissionError().Error())
	}

	// List roles
	roles, pagination, err := r.roleService.List(ctx)

	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusBadRequest, "Role Listing Failed", err.Error())
	}

	return helpers.SuccessResponse(ctx, http.StatusOK, "Roles Retrieved Successfully", http.Json{
		"data":       roles,
		"pagination": pagination,
	});
}

func (r *RoleController) Show(ctx http.Context) http.Response {
	if !helpers.HasPermission(ctx, "view_role") {
		return helpers.ErrorResponse(ctx, http.StatusForbidden, "Validation Error", helpers.NoPermissionError().Error())
	}

	// Show role
	role, err := r.roleService.Show(ctx)

	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusNotFound, "Role not found", err.Error())
	}

	return helpers.SuccessResponse(ctx, http.StatusOK, "Role Retrieved Successfully", helpers.ModelToMap(role))
}


func (r *RoleController) Create(ctx http.Context) http.Response {
	var roleCreateRequest request.RoleCreateRequest

	// Validate request data
	validationErrors, err := ctx.Request().ValidateRequest(&roleCreateRequest)
	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Validation Error", err.Error())
	}

	if validationErrors != nil {
		return helpers.ErrorValidationResponse(ctx, http.StatusUnprocessableEntity, "Validation Error", validationErrors.All())
	}

	// Create role
	role, err := r.roleService.Create(ctx, roleCreateRequest)

	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusBadRequest, "Role Creation Failed", err.Error())
	}

	return helpers.SuccessResponse(ctx, http.StatusCreated, "Role Created Successfully", helpers.ModelToMap(role))
}

func (r *RoleController) Update(ctx http.Context) http.Response {
	var roleUpdateRequest request.RoleUpdateRequest

	// Validate request data
	validationErrors, err := ctx.Request().ValidateRequest(&roleUpdateRequest)
	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Validation Error", err.Error())
	}

	if validationErrors != nil {
		return helpers.ErrorValidationResponse(ctx, http.StatusUnprocessableEntity, "Validation Error", validationErrors.All())
	}

	// Update role
	role, err := r.roleService.Update(ctx, roleUpdateRequest)

	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusNotFound, "Role not found", err.Error())
	}

	return helpers.SuccessResponse(ctx, http.StatusOK, "Role Updated Successfully", helpers.ModelToMap(role))
}

func (r *RoleController) Delete(ctx http.Context) http.Response {
	// Check permission
	if (!helpers.HasPermission(ctx, "delete_role")) {
		return helpers.ErrorResponse(ctx, http.StatusForbidden, "Validation Error", helpers.NoPermissionError().Error())
	}

	// Delete role
	err := r.roleService.Delete(ctx)

	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusNotFound, "Role Deletion Failed", err.Error())
	}

	return helpers.SuccessResponse(ctx, http.StatusOK, "Role Deleted Successfully", nil)
}
