package redis

import (
	"echofy_backend/src/core/utils"
	"fmt"
)

func getAddress() string {
	host := utils.GetenvWithDefault("REDIS_HOST", "redis")
	port := utils.GetenvWithDefault("REDIS_PORT", "6379")

	return fmt.Sprintf("%s:%s", host, port)
}
