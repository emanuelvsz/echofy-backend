package spotifyrepo

import (
	"context"
	"echofy_backend/src/core/domain/authorization"
	"echofy_backend/src/core/errors"
	"echofy_backend/src/core/interfaces/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

var _ repository.AuthLoader = &AuthSpotifyRepository{}

type AuthSpotifyRepository struct{}

var (
	oauth2Config = oauth2.Config{
		ClientID:     "3348f4b437614e8b9c742c305eb9865b",
		ClientSecret: "cd3d51d52a724d18aac6d3910534420c",
		RedirectURL:  "https://5120-186-235-137-201.ngrok-free.app/api/anonymous/callback",
		Scopes:       []string{"user-top-read", "user-library-read", "user-read-playback-state", "user-read-playback-position", "user-read-recently-played", "user-read-currently-playing"},
		Endpoint:     spotify.Endpoint,
	}
)

func (instance *AuthSpotifyRepository) Login(state uuid.UUID) (string, errors.Error) {
	authorizationURL := oauth2Config.AuthCodeURL(state.String())

	return authorizationURL, nil
}

func (instance *AuthSpotifyRepository) Callback(code string) (authorization.Authorization, errors.Error) {

	token, excErr := oauth2Config.Exchange(context.Background(), code)
	if excErr != nil {
		err := errors.NewValidationError("Erro ao obter o Token de Acesso: " + excErr.Error())
		return nil, err
	}

	if token == nil {
		err := errors.NewValidationError("Token de Acesso nulo.")
		return nil, err
	}

	if !token.Valid() {
		err := errors.NewValidationError("Erro durante a troca do código pelo Token de Acesso.")
		return nil, err
	}

	authorization, err := authorization.New(token.AccessToken, token.RefreshToken, time.Since(token.Expiry))
	if err != nil {
		return nil, err
	}

	return authorization, nil
}

func NewAuthSpotifyRepository() *AuthSpotifyRepository {
	return &AuthSpotifyRepository{}
}
