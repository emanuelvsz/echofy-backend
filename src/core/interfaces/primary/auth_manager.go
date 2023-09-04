package primary

import (
	"echofy_backend/src/core/domain/user"
	"echofy_backend/src/core/domain/user/credentials"
	"echofy_backend/src/core/errors"

	"github.com/google/uuid"
)

type AuthManager interface {
	Login(credentials credentials.Credentials) (*user.User, errors.Error)
	SessionExists(userID uuid.UUID, token string) (bool, errors.Error)
}
