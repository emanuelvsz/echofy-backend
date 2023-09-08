package handlers

import (
	"echofy_backend/src/utils/state"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

var (
	oauth2Config = oauth2.Config{
		ClientID:     "3348f4b437614e8b9c742c305eb9865b",
		ClientSecret: "cd3d51d52a724d18aac6d3910534420c",
		RedirectURL:  "https://b0be-201-182-186-221.ngrok-free.app/callback",
		Scopes:       []string{"user-top-read", "user-library-read", "user-read-playback-state", "user-read-playback-position", "user-read-recently-played", "user-read-currently-playing"}, // Adicione as permissões necessárias aqui
		Endpoint:     spotify.Endpoint,
	}

	stateStore = make(map[string]bool)
)

func Authorization(c echo.Context) error {
	state := state.GenerateRandomState()
	authorizationURL := oauth2Config.AuthCodeURL(state)
	stateStore[state] = true
	return c.Redirect(http.StatusFound, authorizationURL)
}
