package middlewares

import (
	"echofy_backend/src/app/api/endpoints/handlers/dtos/response"
	"echofy_backend/src/core/messages"
	"echofy_backend/src/utils/strloader"
	"net/http"
)

var (
	unauthorizedErrorMessage = response.ErrorMessage{
		StatusCode: http.StatusUnauthorized,
		Message:    strloader.Fetch(messages.UnauthorizedErrorMessage),
	}
	forbiddenErrorMessage = response.ErrorMessage{
		StatusCode: http.StatusForbidden,
		Message:    strloader.Fetch(messages.ForbiddenErrorMessage),
	}
)
