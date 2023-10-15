package response

import "echofy_backend/src/core/domain/authorization"

type LoginRedirect struct {
	RedirectURL string `json:"redirect_url"`
}

func NewLoginRedirect(redirectURL string) *LoginRedirect {
	return &LoginRedirect{
		RedirectURL: redirectURL,
	}
}

type Authorization struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewAuthorization(authorization authorization.Authorization) *Authorization {
	return &Authorization{
		Token:        authorization.Token(),
		RefreshToken: authorization.RefreshToken(),
	}
}
