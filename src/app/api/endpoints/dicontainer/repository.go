package dicontainer

import (
	"echofy_backend/src/core/interfaces/repository"
	"echofy_backend/src/infra/postgres"
	spotifyrepo "echofy_backend/src/infra/spotify"
)

func GetAuthRepository() repository.AuthLoader {
	return postgres.NewAuthPostgresRepository(GetPsqlConnectionManager())
}

func GetUserRepository() repository.UserLoader {
	return spotifyrepo.NewUserSpotifyRepository()
}

func GetPsqlConnectionManager() *postgres.DatabaseConnectionManager {
	return &postgres.DatabaseConnectionManager{}
}
