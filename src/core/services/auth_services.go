package services

import (
	userPkg "echofy_backend/src/core/domain/user"
	"echofy_backend/src/core/domain/user/credentials"
	"echofy_backend/src/core/errors"
	"echofy_backend/src/core/errors/logger"
	"echofy_backend/src/core/interfaces/primary"
	"echofy_backend/src/core/interfaces/repository"

	"github.com/google/uuid"
)

var _ primary.AuthManager = (*AuthServices)(nil)

type AuthServices struct {
	authRepository repository.AuthLoader
	logger         logger.Logger
}

func (a *AuthServices) Login(credentials credentials.Credentials) (*userPkg.User, errors.Error) {
	return nil, nil
}

func (a *AuthServices) SessionExists(userID uuid.UUID, token string) (bool, errors.Error) {
	return true, nil
}

func NewAuthServices(authRepository repository.AuthLoader, logger logger.Logger) *AuthServices {
	return &AuthServices{
		authRepository: authRepository,
		logger:         logger,
	}
}