package v1

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	systemMiddleware "github.com/goravel/framework/http/middleware"
	roleController "karuhundeveloper.com/gostarterkit/app/http/controllers/v1/role"
	"karuhundeveloper.com/gostarterkit/app/http/middleware"

	roleService "karuhundeveloper.com/gostarterkit/app/services/v1/role"
)

func V1Role() {
	// Services
	roleService := roleService.NewRoleService()

	// Controllers
	roleController := roleController.NewRoleController(roleService)

	// Authenticated routes
	facades.Route().Middleware(
		systemMiddleware.Throttle("api"),
		middleware.AuthJwtMiddleware(),
	).Prefix("api/v1/roles").Group(func (router route.Router) {
		router.Post("/", roleController.Create).Name("api.v1.roles.create")
		router.Put("/{id}", roleController.Update).Name("api.v1.roles.update")
		router.Delete("/{id}", roleController.Delete).Name("api.v1.roles.delete")
	})
}