package repository

import (
	"echofy_backend/src/core/domain/user"
	"echofy_backend/src/core/domain/user/credentials"
	"echofy_backend/src/core/errors"
)

type AuthLoader interface {
	Login(credentials credentials.Credentials) (*user.User, errors.Error)
}
