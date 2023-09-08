package repository

import (
	"echofy_backend/src/core/errors"
)

type AuthLoader interface {
	Login() errors.Error
}
