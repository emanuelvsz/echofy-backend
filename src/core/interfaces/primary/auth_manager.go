package primary

import (
	"echofy_backend/src/core/errors"
)

type AuthManager interface {
	Login() errors.Error
}
