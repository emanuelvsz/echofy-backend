package spotifyrepo

import (
	"context"
	"echofy_backend/src/core/domain/authorization"
	"echofy_backend/src/core/errors"
	"echofy_backend/src/core/interfaces/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

var _ repository.AuthLoader = &AuthSpotifyRepository{}

type AuthSpotifyRepository struct {
	oauth2Config *oauth2.Config
}

func (instance *AuthSpotifyRepository) Login(state uuid.UUID) (string, errors.Error) {
	authorizationURL := instance.oauth2Config.AuthCodeURL(state.String())

	return authorizationURL, nil
}

func (instance *AuthSpotifyRepository) Callback(code string) (authorization.Authorization, errors.Error) {

	token, excErr := instance.oauth2Config.Exchange(context.Background(), code)
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

func NewAuthSpotifyRepository(oauth2Config oauth2.Config) *AuthSpotifyRepository {
	return &AuthSpotifyRepository{
		oauth2Config: &oauth2Config,
	}
}
