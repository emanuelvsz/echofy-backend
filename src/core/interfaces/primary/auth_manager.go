package primary

import (
	"echofy_backend/src/core/errors"

	"github.com/google/uuid"
)

type AuthManager interface {
	Login() (string, errors.Error)
	Callback(code string, state uuid.UUID) errors.Error
}
