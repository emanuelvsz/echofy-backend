package dicontainer

import (
	"echofy_backend/src/core/interfaces/repository"
	"echofy_backend/src/infra/postgres"
)

func GetAuthRepository() repository.AuthLoader {
	return postgres.NewAuthPostgresRepository(GetPsqlConnectionManager())
}

func GetUserRepository() repository.UserLoader {
	return postgres.NewUserPostgresRepository(GetPsqlConnectionManager())
}

func GetPsqlConnectionManager() *postgres.DatabaseConnectionManager {
	return &postgres.DatabaseConnectionManager{}
}
