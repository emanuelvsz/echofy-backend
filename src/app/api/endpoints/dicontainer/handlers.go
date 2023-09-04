package dicontainer

import "echofy_backend/src/app/api/endpoints/handlers"

func GetAuthHandlers() *handlers.AuthHandlers {
	return handlers.NewAuthHandlers(GetAuthServices())
}

func GetUserHandlers() *handlers.UserHandlers {
	return handlers.NewUserHandlers(GetUserServices())
}
