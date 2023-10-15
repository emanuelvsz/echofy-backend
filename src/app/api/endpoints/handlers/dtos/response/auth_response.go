package response

type LoginRedirect struct {
	RedirectURL string `json:"redirect_url"`
}

func NewLoginRedirect(redirectURL string) *LoginRedirect {
	return &LoginRedirect{
		RedirectURL: redirectURL,
	}
}
