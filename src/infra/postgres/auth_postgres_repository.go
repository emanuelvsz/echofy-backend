package postgres

import (
	"echofy_backend/src/core/errors"
)

// var _ repository.AuthLoader = &AuthPostgresRepository{}

type AuthPostgresRepository struct {
	connectorManager
}

func (instance *AuthPostgresRepository) Login() errors.Error {
	return nil
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
