package dicontainer

import (
	"echofy_backend/src/core/errors/logger"
	"echofy_backend/src/core/interfaces/primary"
	"echofy_backend/src/core/services"
)

func GetAuthServices() primary.AuthManager {
	return services.NewAuthServices(GetAuthRepository(), GetSessionRepository(), GetLogger())
}

func GetUserServices() primary.UserManager {
	return services.NewUserServices(GetUserRepository(), GetLogger())
}

func GetLogger() logger.Logger {
	return logger.New()
}
