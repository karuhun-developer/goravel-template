package middleware

import (
	"errors"

	"github.com/goravel/framework/auth"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func AuthJwtMiddleware() http.Middleware {
	return func(ctx http.Context) {
		// Check jwt token validity here
		originalToken := ctx.Request().Header("Authorization")

		// If no token is provided, abort with unauthorized status
		if originalToken == "" {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
				"message": "Unauthorized: No token provided",
			})
			return
		}

		// Validate the token (this is a placeholder, implement your own logic)
		_, err := facades.Auth(ctx).Parse(originalToken)

		// If token is invalid, abort with unauthorized status
		if err != nil {
			// If error token expired
			if errors.Is(err, auth.ErrorTokenExpired) {
				newToken, err := facades.Auth(ctx).Refresh()

				if err != nil {
					ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
						"message": "Unauthorized: Token expired and refresh failed",
					})
					return
				}

				// Set new token in the response header
				originalToken = "Bearer " + newToken
			} else {
				ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
					"message": "Unauthorized: Invalid token",
				})
				return
			}
		}

		// Token is valid, proceed to the next middleware/handler
		ctx.Request().Header("Authorization", originalToken)
		ctx.Request().Next()
	}
}
