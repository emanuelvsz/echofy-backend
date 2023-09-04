package user

import (
	"echofy_backend/src/core"
	"echofy_backend/src/core/errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var logger = core.Logger()

const (
	TokenTimeout      = time.Hour
	AnonymousRoleCode = "anonymous"
)

type Authorization interface {
	Token() string
}

type authorization struct {
	token string
}

func New() Authorization {
	return &authorization{}
}

func NewFromUser(user User) (Authorization, errors.Error) {
	instance := &authorization{}
	if err := instance.GenerateToken(user); err != nil {
		return nil, err
	}
	return instance, nil
}

func (a *authorization) Token() string {
	return a.token
}

func (a *authorization) GenerateToken(user User) errors.Error {
	secret := os.Getenv("SERVER_SECRET")
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims(
		user.ID().String(),
		"bearer",
		time.Now().Add(TokenTimeout).Unix(),
	)).SignedString([]byte(secret))
	if err != nil {
		logger.Error().Msg(err.Error())
		return errors.NewUnexpectedError(err.Error(), err)
	}
	a.token = token
	return nil
}
