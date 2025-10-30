package v1

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	systemMiddleware "github.com/goravel/framework/http/middleware"
	"karuhundeveloper.com/gostarterkit/app/http/middleware"
)

func V1Auth() {
	// Guest routes
	facades.Route().Middleware(
		systemMiddleware.Throttle("api"),
		middleware.AuthJwtMiddleware(),
	).Prefix("api/v1").Group(func (router route.Router) {
		// Define authenticated routes here
	})
}