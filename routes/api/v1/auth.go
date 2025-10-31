package v1

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	systemMiddleware "github.com/goravel/framework/http/middleware"
	authController "karuhundeveloper.com/gostarterkit/app/http/controllers/v1/auth"
	"karuhundeveloper.com/gostarterkit/app/http/middleware"

	authService "karuhundeveloper.com/gostarterkit/app/services/v1/auth"
)

func V1Auth() {
	// Services
	authService := authService.NewAuthService()

	// Controllers
	authController := authController.NewAuthController(authService)

	// Guest routes
	facades.Route().Middleware(
		systemMiddleware.Throttle("api"),
	).Prefix("api/v1/auth").Group(func (router route.Router) {
		// Auth routes
		router.Post("login", authController.Login).Name("api.v1.auth.login")

	})

	// Authenticated routes
	facades.Route().Middleware(
		systemMiddleware.Throttle("api"),
		middleware.AuthJwtMiddleware(),
	).Prefix("api/v1/auth").Group(func (router route.Router) {
		// Auth routes
		router.Put("refresh-token", authController.RefreshToken).Name("api.v1.auth.refresh-token")
		router.Post("logout", authController.Logout).Name("api.v1.auth.logout")
	})
}