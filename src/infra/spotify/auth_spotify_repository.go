package spotifyrepo

import (
	"echofy_backend/src/core/errors"
	"echofy_backend/src/core/interfaces/repository"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

var _ repository.AuthLoader = &AuthSpotifyRepository{}

type AuthSpotifyRepository struct{}

var (
	oauth2Config = oauth2.Config{
		ClientID:     "3348f4b437614e8b9c742c305eb9865b",
		ClientSecret: "cd3d51d52a724d18aac6d3910534420c",
		RedirectURL:  "https://b0be-201-182-186-221.ngrok-free.app/callback",
		Scopes:       []string{"user-top-read", "user-library-read", "user-read-playback-state", "user-read-playback-position", "user-read-recently-played", "user-read-currently-playing"},
		Endpoint:     spotify.Endpoint,
	}

	stateStore = make(map[string]bool)
)

func (instance *AuthSpotifyRepository) Login() errors.Error {
	
	return nil
}

func NewAuthSpotifyRepository() *AuthSpotifyRepository {
	return &AuthSpotifyRepository{}
}
