package services

import (
	"echofy_backend/src/core/errors"
	"echofy_backend/src/core/errors/logger"
	"echofy_backend/src/core/interfaces/primary"
	"echofy_backend/src/core/interfaces/repository"
)

var _ primary.AuthManager = (*AuthServices)(nil)

type AuthServices struct {
	authRepository repository.AuthLoader
	logger         logger.Logger
}

func (a *AuthServices) Login() errors.Error {
	return nil
}

func NewAuthServices(authRepository repository.AuthLoader, logger logger.Logger) *AuthServices {
	return &AuthServices{
		authRepository: authRepository,
		logger:         logger,
	}
}
