package repository

import (
	"echofy_backend/src/core/domain/authorization"
	"echofy_backend/src/core/errors"

	"github.com/google/uuid"
)

type SessionLoader interface {
	CreateSession(key uuid.UUID) errors.Error
	StoreAuth(key uuid.UUID, auth authorization.Authorization) errors.Error
	StateExists(key uuid.UUID) (bool, errors.Error)
	Close(key uuid.UUID) errors.Error
}
