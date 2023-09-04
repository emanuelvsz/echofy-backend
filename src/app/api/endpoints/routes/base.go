package routes

import (
	_ "echofy_backend/src/app/api/docs"
	"echofy_backend/src/app/api/endpoints/routes/users"
	"github.com/labstack/echo/v4"
)

type Router interface {
	Load(*echo.Group)
}

type router struct {
}

func New() Router {
	return &router{}
}

func (instance *router) Load(group *echo.Group) {
	loadApiRoutes(group)
	users.LoadUserRoutes(group)
}
