package user

import (
	"echofy_backend/src/app/api/endpoints/dicontainer"

	"github.com/labstack/echo/v4"
)

var userHandlers = dicontainer.GetUserHandlers()

func LoadUserRoutes(group *echo.Group) {
	userGroup := group.Group("/user")
	userHandlers := dicontainer.GetUserHandlers()

	userGroup.GET("", userHandlers.GetUserBasicInfo)
	userGroup.GET("/playlist/:playlistID", userHandlers.GetPlaylistByID)
	userGroup.GET("/playlist/:playlistID/songs", userHandlers.GetSongsByPlaylistID)
	userGroup.GET("/album/:albumID/songs", userHandlers.GetAlbumTracks)
	userGroup.GET("/artist/:artistID/albums", userHandlers.GetArtistAlbums)
	userGroup.GET("/song/:songID/details", userHandlers.FetchSongDetailsByID)
}
