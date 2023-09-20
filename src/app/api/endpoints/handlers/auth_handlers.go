package handlers

import (
	"echofy_backend/src/core/interfaces/primary"
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
// @Router /anonymous/authenticate [get]
func (instance AuthHandlers) Login(context echo.Context) error {
	return context.JSON(http.StatusCreated, nil)
}

func NewAuthHandlers(service primary.AuthManager) *AuthHandlers {
	return &AuthHandlers{service: service}
}
