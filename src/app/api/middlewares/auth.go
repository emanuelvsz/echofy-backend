package middlewares

import (
	"echofy_backend/src/app/api/endpoints/handlers/dtos/response"
	"echofy_backend/src/app/api/utils"
	"echofy_backend/src/core/messages"
	"echofy_backend/src/utils/strloader"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
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

func GuardMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	authModel := os.Getenv("SERVER_CASBIN_AUTH_MODEL")
	authPolicy := os.Getenv("SERVER_CASBIN_AUTH_POLICY")

	enforcer, err := casbin.NewEnforcer(authModel, authPolicy)
	if err != nil {
		fmt.Println("Error when building enforcer:", err)
		log.Fatal().Err(err)
	}

	return func(context echo.Context) error {
		authHeader := context.Request().Header.Get("Authorization")
		method := context.Request().Method
		path := context.Request().URL.Path

		authType, authToken := utils.ExtractToken(authHeader)
		if strings.TrimSpace(authToken) == "" {
			if ok, err := enforcer.Enforce("anonymous", path, method); err != nil {
				fmt.Println("Error when doing enforce:", err)
				log.Fatal().Err(err)
			} else if !ok {
				return context.JSON(unauthorizedErrorMessage.StatusCode, unauthorizedErrorMessage)
			}

			return next(context)
		}

		claims := utils.ExtractAuthorizationClaims(authType, authToken)
		if claims == nil {
			return context.JSON(forbiddenErrorMessage.StatusCode, forbiddenErrorMessage)
		}
		return next(context)
	}
}
