package responses

import "github.com/goravel/framework/contracts/http"

func ErrorResponse(ctx http.Context, status int, message string, err string) http.Response {
	return ctx.Response().Json(status, http.Json{
		"code":		status,
		"message": 	message,
		"data":		err,
	})
}

func ErrorValidationResponse(ctx http.Context, status int, message string, errors map[string]map[string]string) http.Response {
	return ctx.Response().Json(status, http.Json{
		"code":		status,
		"message": 	message,
		"data":		errors,
	})
}

func SuccessResponse(ctx http.Context, status int, message string, data http.Json) http.Response {
	return ctx.Response().Json(status, http.Json{
		"code":		status,
		"message": 	message,
		"data":		data,
	})
}