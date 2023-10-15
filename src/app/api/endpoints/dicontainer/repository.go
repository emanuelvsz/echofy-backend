package dicontainer

import (
	"echofy_backend/src/core/interfaces/repository"
	"echofy_backend/src/infra/postgres"
	"echofy_backend/src/infra/redis"
	spotifyrepo "echofy_backend/src/infra/spotify"
)

func GetAuthRepository() repository.AuthLoader {
	return spotifyrepo.NewAuthSpotifyRepository()
}

func GetSessionRepository() repository.SessionLoader {
	return redis.NewSessionRepository(redis.RedisConnectionManager{})
}

func GetUserRepository() repository.UserLoader {
	return spotifyrepo.NewUserSpotifyRepository()
}

func GetPsqlConnectionManager() *postgres.DatabaseConnectionManager {
	return &postgres.DatabaseConnectionManager{}
}
