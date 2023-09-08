package handlers

import (
	"echofy_backend/src/core/interfaces/primary"
	"echofy_backend/src/core/messages"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandlers struct {
	service primary.AuthManager
}

// Login
// @ID Login
// @Summary Fazer a autenticação no sistema
// @Tags Rotas anônimas
// @Description Rota que permite que um usuário se autentique no Echofy com seus dados de sua conta do Spotify.										  |
// @Accept json
// @Produce json
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /anonymous/login [get]
func (instance AuthHandlers) Login(context echo.Context) error {
	return context.JSON(http.StatusCreated, nil)
}

// CallBack
// @ID CallBack
// @Summary Fazer a autenticação no sistema
// @Tags Rotas anônimas
// @Description Esta rota é usada para processar o código de autorização e o estado.
// @Accept json
// @Produce json
// @Param code query string false "O código de autorização recebido após a autenticação."
// @Param state query string false "Um valor de estado aleatório usado para proteção CSRF"
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /anonymous/callback [get]
func (instance AuthHandlers) CallBack(context echo.Context) error {
	code := context.QueryParam("code")
	state := context.QueryParam("state")

	if !stateStore[state] {
		return context.JSON(http.StatusBadRequest, messages.CallBackStateInvalidErrMsg)
	}

	token, err := oauth2Config.Exchange(context.Request().Context(), code)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, messages.ErrorWhileObtainAccessToken)
	}

	if token == nil {
		return context.JSON(http.StatusInternalServerError, messages.NullAccessToken)
	}

	if !token.Valid() {
		return context.JSON(http.StatusInternalServerError, messages.ErrorWhileTradeCodeForAccessToken)
	}

	return context.JSON(http.StatusOK, token)
}

func NewAuthHandlers(service primary.AuthManager) *AuthHandlers {
	return &AuthHandlers{service: service}
}
