package response

import (
	"echofy_backend/src/core/domain/user"
)

type Authorization struct {
	Token string `json:"access_token"`
}

func NewAuthorization(authorization user.Authorization) *Authorization {
	return &Authorization{Token: authorization.Token()}
}