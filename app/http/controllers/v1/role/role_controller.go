package role

import (
	"github.com/goravel/framework/contracts/http"
	request "karuhundeveloper.com/gostarterkit/app/http/requests/v1/role"
	"karuhundeveloper.com/gostarterkit/app/http/responses"
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

func (r *RoleController) Create(ctx http.Context) http.Response {
	var roleCreateRequest request.RoleCreateRequest

	// Validate request data
	validationErrors, err := ctx.Request().ValidateRequest(&roleCreateRequest)
	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Validation Error", err.Error())
	}

	if validationErrors != nil {
		return responses.ErrorValidationResponse(ctx, http.StatusUnprocessableEntity, "Validation Error", validationErrors.All())
	}

	// Create role
	role, err := r.roleService.Create(ctx, roleCreateRequest)

	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusInternalServerError, "Role Creation Failed", err.Error())
	}

	return responses.SuccessResponse(ctx, http.StatusCreated, "Role Created Successfully", http.Json{
		"id":   role.ID,
		"name": role.Name,
		"created_at": role.CreatedAt,
		"updated_at": role.UpdatedAt,
	})
}

func (r *RoleController) Update(ctx http.Context) http.Response {
	var roleUpdateRequest request.RoleUpdateRequest

	// Validate request data
	validationErrors, err := ctx.Request().ValidateRequest(&roleUpdateRequest)
	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Validation Error", err.Error())
	}

	if validationErrors != nil {
		return responses.ErrorValidationResponse(ctx, http.StatusUnprocessableEntity, "Validation Error", validationErrors.All())
	}

	// Update role
	role, err := r.roleService.Update(ctx, roleUpdateRequest)

	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusInternalServerError, "Role Update Failed", err.Error())
	}

	return responses.SuccessResponse(ctx, http.StatusOK, "Role Updated Successfully", http.Json{
		"id":   role.ID,
		"name": role.Name,
		"created_at": role.CreatedAt,
		"updated_at": role.UpdatedAt,
	})
}

func (r *RoleController) Delete(ctx http.Context) http.Response {
	// Delete role
	err := r.roleService.Delete(ctx)

	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusInternalServerError, "Role Deletion Failed", err.Error())
	}

	return responses.SuccessResponse(ctx, http.StatusOK, "Role Deleted Successfully", nil)
}
