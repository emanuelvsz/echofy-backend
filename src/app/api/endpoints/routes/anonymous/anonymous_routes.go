package anonymous

import (
	"echofy_backend/src/app/api/endpoints/dicontainer"

	"github.com/labstack/echo/v4"
)

func LoadAnonymousRoute(group *echo.Group) {
	anonyousGroup := group.Group("/anonymous")
	userHandlers := dicontainer.GetUserHandlers()

	anonyousGroup.GET("/authorize", userHandlers.Authorize)
}
