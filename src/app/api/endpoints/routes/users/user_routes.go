package users

import (
	"echofy_backend/src/app/api/endpoints/dicontainer"

	"github.com/labstack/echo/v4"
)

var userHandlers = dicontainer.GetUserHandlers()

func LoadUserRoutes(group *echo.Group) {

}
