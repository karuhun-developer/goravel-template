package auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/errors"

	"github.com/goravel/framework/facades"
	"karuhundeveloper.com/gostarterkit/app/http/requests/auth"
	"karuhundeveloper.com/gostarterkit/app/models/user"
)

type AuthService struct {
	// Add your service dependencies here
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (u *AuthService) Login(ctx http.Context, loginRequest auth.LoginRequest) (token string, user user.User, err error) {
	// Check is user exists
	err = facades.Orm().Query().
		Where("email", loginRequest.Email).
		With("Roles").
		FirstOrFail(&user)

	// Return if user not found
	if (errors.Is(err, errors.OrmRecordNotFound)) {
		err = errors.New("invalid credentials")
		return
	}

	// Verify password
	if !facades.Hash().Check(loginRequest.Password, user.Password) {
		err = errors.New("invalid credentials")
		return
	}

	// Generate token
	token, err = facades.Auth(ctx).Login(&user)
	if err != nil {
		return
	}

	return
}

func (u *AuthService) Logout(ctx http.Context) error {
	return facades.Auth(ctx).Logout()
}