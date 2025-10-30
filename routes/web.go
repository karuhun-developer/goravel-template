package routes

import (
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"status":      "OK",
			"server_time": time.Now().Format(time.RFC850), // includes zone name and offset
			"version":     support.Version,
		})
	})
}
