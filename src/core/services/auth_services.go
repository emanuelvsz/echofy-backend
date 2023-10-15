package services

import (
	"echofy_backend/src/core/errors"
	"echofy_backend/src/core/errors/logger"
	"echofy_backend/src/core/interfaces/primary"
	"echofy_backend/src/core/interfaces/repository"

	"github.com/google/uuid"
)

var _ primary.AuthManager = (*AuthServices)(nil)

type AuthServices struct {
	authRepository    repository.AuthLoader
	sessionRepository repository.SessionLoader
	logger            logger.Logger
}

func (a *AuthServices) Login() (string, errors.Error) {
	state := uuid.New()

	authURL, err := a.authRepository.Login(state)
	if err != nil {
		a.logger.Log(err)
		return "", err
	}

	err = a.sessionRepository.CreateSession(state)
	if err != nil {
		a.logger.Log(err)
		return "", err
	}

	return authURL, nil
}

func (a *AuthServices) Callback(code string, state uuid.UUID) errors.Error {
	sessionExists, err := a.sessionRepository.StateExists(state)
	if err != nil {
		a.logger.Log(err)
		return err
	}

	if !sessionExists {
		err := errors.NewValidationError("A sessão não existe!")
		a.logger.Log(err)
		return err
	}

	authorization, err := a.authRepository.Callback(code)
	if err != nil {
		a.logger.Log(err)
		return err
	}

	err = a.sessionRepository.StoreAuth(state, authorization)
	if err != nil {
		a.logger.Log(err)
		return err
	}

	return nil
}

func NewAuthServices(authRepository repository.AuthLoader, sessionRepository repository.SessionLoader, logger logger.Logger) *AuthServices {
	return &AuthServices{
		authRepository:    authRepository,
		sessionRepository: sessionRepository,
		logger:            logger,
	}
}
