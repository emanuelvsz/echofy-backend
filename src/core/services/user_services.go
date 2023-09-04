package services

import (
	"echofy_backend/src/core/errors/logger"
	"echofy_backend/src/core/interfaces/primary"
	"echofy_backend/src/core/interfaces/repository"
)

var _ primary.UserManager = (*UserServices)(nil)

type UserServices struct {
	userRepository repository.UserLoader
	logger         logger.Logger
}

func NewUserServices(userRepository repository.UserLoader, logger logger.Logger) *UserServices {
	return &UserServices{
		userRepository: userRepository,
		logger:         logger,
	}
}
