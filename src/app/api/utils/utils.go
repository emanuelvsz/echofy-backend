package utils

import (
	"echofy_backend/src/core"
	"echofy_backend/src/core/domain/user"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

var logger = core.Logger()

func ExtractToken(authHeader string) (authType string, token string) {
	authorization := strings.Split(strings.TrimSpace(authHeader), " ")
	if len(authorization) < 2 {
		return "", ""
	}
	authType = authorization[0]
	token = authorization[1]
	return authType, token
}

func ExtractAuthorizationClaims(authType, authToken string) *user.Claims {
	secret := os.Getenv("SERVER_SECRET")
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		logger.Error().Msg("error parsing the provided token on (signature is invalid?)")
		return nil
	}
	if !token.Valid || token.Claims.Valid() != nil {
		logger.Error().Msg("the provided token is invalid or expired")
		return nil
	}
	claims, err := extractTokenClaims(authToken)
	if err != nil || claims == nil {
		return nil
	}
	if !strings.EqualFold(claims.Type, authType) {
		logger.Error().Msg(fmt.Sprintf("the used authorization type \"%s\" is not  \"%s\"", authType, claims.Type))
		return nil
	}
	return claims
}

func extractTokenClaims(authToken string) (*user.Claims, error) {
	parts := strings.Split(authToken, ".")
	if len(parts) != 3 {
		return nil, errors.New("Invalid token, it must to be with 3 parts!")
	}
	payload := parts[1]
	payloadBytes, err := jwt.DecodeSegment(payload)
	if err != nil {
		logger.Error().Msg("an error occurred when decoding the token payload: " + err.Error())
		return nil, err
	}

	var claims user.Claims
	err = json.Unmarshal(payloadBytes, &claims)
	if err != nil {
		logger.Error().Msg("an error occurred when unmarshalling the token payload: " + err.Error())
		return nil, err
	}
	return &claims, nil
}
