package role

import (
	"github.com/goravel/framework/contracts/http"
	"karuhundeveloper.com/gostarterkit/app/helpers"
	request "karuhundeveloper.com/gostarterkit/app/http/requests/v1/role"
	"karuhundeveloper.com/gostarterkit/app/services/v1/role"
)

type PermissionController struct {
	// Dependent services
	permissionService *role.PermissionService
}

func NewPermissionController(PermissionService *role.PermissionService) *PermissionController {
	return &PermissionController{
		permissionService: PermissionService,
	}
}

func (r *PermissionController) Index(ctx http.Context) http.Response {
	if !helpers.HasPermission(ctx, "view_all_permission") {
		return helpers.ErrorResponse(ctx, http.StatusForbidden, "Validation Error", helpers.NoPermissionError().Error())
	}

	// List permissions
	_, pagination, err := r.permissionService.List(ctx)

	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusBadRequest, "Permission Listing Failed", err.Error())
	}

	return helpers.SuccessResponse(ctx, http.StatusOK, "Permissions Retrieved Successfully", http.Json{
		"pagination": pagination,
	});
}

func (r *PermissionController) Show(ctx http.Context) http.Response {
	if !helpers.HasPermission(ctx, "view_permission") {
		return helpers.ErrorResponse(ctx, http.StatusForbidden, "Validation Error", helpers.NoPermissionError().Error())
	}

	// Show permission
	permission, err := r.permissionService.Show(ctx)

	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusNotFound, "Permission not found", err.Error())
	}

	return helpers.SuccessResponse(ctx, http.StatusOK, "Permission Retrieved Successfully", helpers.ModelToMap(permission))
}


func (r *PermissionController) Create(ctx http.Context) http.Response {
	var permissionCreateRequest request.PermissionCreateRequest

	// Validate request data
	validationErrors, err := ctx.Request().ValidateRequest(&permissionCreateRequest)
	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Validation Error", err.Error())
	}

	if validationErrors != nil {
		return helpers.ErrorValidationResponse(ctx, http.StatusUnprocessableEntity, "Validation Error", validationErrors.All())
	}

	// Create permission
	permission, err := r.permissionService.Create(ctx, permissionCreateRequest)

	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusBadRequest, "Permission Creation Failed", err.Error())
	}

	return helpers.SuccessResponse(ctx, http.StatusCreated, "Permission Created Successfully", helpers.ModelToMap(permission))
}

func (r *PermissionController) Update(ctx http.Context) http.Response {
	var permissionUpdateRequest request.PermissionUpdateRequest

	// Validate request data
	validationErrors, err := ctx.Request().ValidateRequest(&permissionUpdateRequest)
	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Validation Error", err.Error())
	}

	if validationErrors != nil {
		return helpers.ErrorValidationResponse(ctx, http.StatusUnprocessableEntity, "Validation Error", validationErrors.All())
	}

	// Update permission
	permission, err := r.permissionService.Update(ctx, permissionUpdateRequest)

	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusNotFound, "Permission not found", err.Error())
	}

	return helpers.SuccessResponse(ctx, http.StatusOK, "Permission Updated Successfully", helpers.ModelToMap(permission))
}

func (r *PermissionController) Delete(ctx http.Context) http.Response {
	// Check permission
	if (!helpers.HasPermission(ctx, "delete_permission")) {
		return helpers.ErrorResponse(ctx, http.StatusForbidden, "Validation Error", helpers.NoPermissionError().Error())
	}

	// Delete permission
	err := r.permissionService.Delete(ctx)

	if err != nil {
		return helpers.ErrorResponse(ctx, http.StatusNotFound, "Permission Deletion Failed", err.Error())
	}

	return helpers.SuccessResponse(ctx, http.StatusOK, "Permission Deleted Successfully", nil)
}
