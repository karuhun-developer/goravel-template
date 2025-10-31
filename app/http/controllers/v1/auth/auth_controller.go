package auth

import (
	"github.com/goravel/framework/contracts/http"
	request "karuhundeveloper.com/gostarterkit/app/http/requests/v1/auth"
	"karuhundeveloper.com/gostarterkit/app/http/responses"
	"karuhundeveloper.com/gostarterkit/app/services/v1/auth"
)

type AuthController struct {
	// Dependent services
	authService *auth.AuthService
}

func NewAuthController(AuthService *auth.AuthService) *AuthController {
	return &AuthController{
		authService: AuthService,
	}
}

func (r *AuthController) Login(ctx http.Context) http.Response {
	var loginRequest request.LoginRequest

	// Validate request data
	validationErrors, err := ctx.Request().ValidateRequest(&loginRequest)
	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Validation Error", err.Error())
	}

	if validationErrors != nil {
		return responses.ErrorValidationResponse(ctx, http.StatusUnprocessableEntity, "Validation Error", validationErrors.All())
	}

	// Get authentication data from request
	token, user, err := r.authService.Login(ctx, loginRequest)

	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusUnauthorized, "Login Failed", err.Error())
	}

	return responses.SuccessResponse(ctx, http.StatusOK, "Login Successful", http.Json{
		"token": token,
		"user":  user,
	})
}

func (r *AuthController) RefreshToken(ctx http.Context) http.Response {
	token, err := r.authService.RefreshToken(ctx)
	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusUnauthorized, "Token Refresh Failed", err.Error())
	}
	return responses.SuccessResponse(ctx, http.StatusOK, "Token Refresh Successful", http.Json{
		"token": token,
	})
}

func (r *AuthController) Logout(ctx http.Context) http.Response {
	err := r.authService.Logout(ctx)
	if err != nil {
		return responses.ErrorResponse(ctx, http.StatusUnauthorized, "Logout Failed", err.Error())
	}
	return responses.SuccessResponse(ctx, http.StatusOK, "Logout Successful", nil)
}
