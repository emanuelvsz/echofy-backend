package postgres

import (
	u "echofy_backend/src/core/domain/user"
	"echofy_backend/src/core/domain/user/credentials"
	"echofy_backend/src/core/errors"
	"echofy_backend/src/core/interfaces/repository"
)

var _ repository.AuthLoader = &AuthPostgresRepository{}

type AuthPostgresRepository struct {
	connectorManager
}

func (instance *AuthPostgresRepository) Login(credentials credentials.Credentials) (*u.User, errors.Error) {
	return nil, nil
}

// func (instance AuthPostgresRepository) handleError(err error) errors.Error {
// 	msg := err.Error()

// 	if strings.Contains(msg, "sql: no rows in result set") {
// 		return errors.NewNotFoundError(messages.InvalidCredentialsErrorMessage, err)
// 	}

// 	return errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
// }

func NewAuthPostgresRepository(manager connectorManager) *AuthPostgresRepository {
	return &AuthPostgresRepository{manager}
}
