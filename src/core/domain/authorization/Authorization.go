package authorization

import (
	"echofy_backend/src/core/errors"
	"time"
)

const (
	RefreshTokenTimeout = time.Hour
	SessionTimeout      = time.Minute * 10
	AnonymousRoleCode   = "anonymous"
)

type Authorization interface {
	Token() string
	RefreshToken() string
	Expiry() time.Duration
}

type authorization struct {
	token        string
	refreshToken string
	expiry       time.Duration
}

func (a *authorization) Token() string {
	return a.token
}

func (a *authorization) RefreshToken() string {
	return a.refreshToken
}

func (a *authorization) Expiry() time.Duration {
	return a.expiry
}

func New(accessToken string, refreshToken string, expiry time.Duration) (Authorization, errors.Error) {
	return &authorization{
		token:        accessToken,
		refreshToken: refreshToken,
		expiry:       expiry,
	}, nil
}
