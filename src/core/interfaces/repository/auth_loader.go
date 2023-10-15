package repository

import (
	"echofy_backend/src/core/domain/authorization"
	"echofy_backend/src/core/errors"

	"github.com/google/uuid"
)

type AuthLoader interface {
	Login(state uuid.UUID) (string, errors.Error)
	Callback(code string) (authorization.Authorization, errors.Error)
}
