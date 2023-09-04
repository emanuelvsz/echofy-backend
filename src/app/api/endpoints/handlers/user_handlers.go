package handlers

import (
	"echofy_backend/src/core/interfaces/primary"
)

const (
	albumID  = "albumID"
	artistID = "artistID"
	songID   = "songID"
)

type UserHandlers struct {
	service primary.UserManager
}

func NewUserHandlers(service primary.UserManager) *UserHandlers {
	return &UserHandlers{service: service}
}
