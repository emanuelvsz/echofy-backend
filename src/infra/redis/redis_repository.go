package redis

import (
	"echofy_backend/src/core/utils"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
)

type connectorManager interface {
	getConnection() (*redis.Client, error)
	closeConnection(conn *redis.Client)
}

var _ connectorManager = (*RedisConnectionManager)(nil)

type RedisConnectionManager struct{}

func (rdm RedisConnectionManager) getConnection() (*redis.Client, error) {
	password := utils.GetenvWithDefault("REDIS_PASSWORD", "")

	options := redis.Options{
		Addr:     getAddress(),
		Password: password,
		DB:       0,
	}
	conn := redis.NewClient(&options)

	if result := conn.Ping(); result.Err() != nil {
		log.Print("Error while accessing redis database: ", result.Err().Error())
		return nil, result.Err()
	}
	return conn, nil
}

func (rcm RedisConnectionManager) closeConnection(conn *redis.Client) {
	err := conn.Close()
	if err != nil {
		log.Error().Err(err)
	}
}
