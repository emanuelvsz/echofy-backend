package users

import (
	"echofy_backend/src/app/api/endpoints/dicontainer"

	"github.com/labstack/echo/v4"
)

var userHandlers = dicontainer.GetUserHandlers()

func LoadUserRoutes(group *echo.Group) {
	userGroup := group.Group("/user")
	userHandlers := dicontainer.GetUserHandlers()

	userGroup.GET("/playlist/:playlistID/songs", userHandlers.GetSongsByPlaylistID)
}
