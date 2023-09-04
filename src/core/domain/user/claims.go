package user

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.Claims `json:"c,emitempty"`
	AccountID  string `json:"sub"`
	Expiry     int64  `json:"exp"`
	Type       string `json:"typ"`
}

func newClaims(accountID string, typ string, exp int64) *Claims {
	return &Claims{
		AccountID: accountID,
		Type:      typ,
		Expiry:    exp,
	}
}
