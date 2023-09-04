package postgres

import (
	"echofy_backend/src/core/interfaces/repository"
)

var _ repository.UserLoader = &UserPostgresRepository{}

type UserPostgresRepository struct {
	connectorManager
}

func NewUserPostgresRepository(manager connectorManager) *UserPostgresRepository {
	return &UserPostgresRepository{manager}
}
