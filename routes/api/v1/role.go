package v1

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	systemMiddleware "github.com/goravel/framework/http/middleware"
	controller "karuhundeveloper.com/gostarterkit/app/http/controllers/v1/role"
	"karuhundeveloper.com/gostarterkit/app/http/middleware"

	service "karuhundeveloper.com/gostarterkit/app/services/v1/role"
)

func V1Role() {
	// Services
	roleService := service.NewRoleService()
	permissionService := service.NewPermissionService()
	rolePermissionService := service.NewRolePermissionService()

	// Controllers
	roleController := controller.NewRoleController(roleService)
	permissionController := controller.NewPermissionController(permissionService)
	rolePermissionController := controller.NewRolePermissionController(rolePermissionService)

	// Authenticated routes
	facades.Route().Middleware(
		systemMiddleware.Throttle("api"),
		middleware.AuthJwtMiddleware(),
	).Prefix("api/v1/role").Group(func (router route.Router) {
		// Role Routes
		router.Get("/roles", roleController.Index).Name("api.v1.roles.index")
		router.Get("/roles/{id}", roleController.Show).Name("api.v1.roles.show")
		router.Post("/roles", roleController.Create).Name("api.v1.roles.create")
		router.Put("/roles/{id}", roleController.Update).Name("api.v1.roles.update")
		router.Delete("/roles/{id}", roleController.Delete).Name("api.v1.roles.delete")

		// Permission Routes
		router.Get("/permissions", permissionController.Index).Name("api.v1.permissions.index")
		router.Get("/permissions/{id}", permissionController.Show).Name("api.v1.permissions.show")
		router.Post("/permissions", permissionController.Create).Name("api.v1.permissions.create")
		router.Put("/permissions/{id}", permissionController.Update).Name("api.v1.permissions.update")
		router.Delete("/permissions/{id}", permissionController.Delete).Name("api.v1.permissions.delete")

		// Role Permission Routes
		router.Get("/roles/{id}/permissions", rolePermissionController.Show).Name("api.v1.roles.permissions.show")
		router.Put("/roles/{id}/permissions", rolePermissionController.Sync).Name("api.v1.roles.permissions.sync")
	})
}