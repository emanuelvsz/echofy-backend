package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func loadApiRoutes(group *echo.Group) {
	group.GET("/*", echoSwagger.WrapHandler)
}
