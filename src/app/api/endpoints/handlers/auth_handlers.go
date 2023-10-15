package handlers

import (
	"echofy_backend/src/app/api/endpoints/handlers/dtos/response"
	"echofy_backend/src/core/errors"
	"echofy_backend/src/core/interfaces/primary"
	"echofy_backend/src/core/utils"
	"net/http"

	"github.com/google/uuid"
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
// @Success 307 {object} response.LoginRedirect "Autorização feita com sucesso! Redirecionando o usuário"
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /anonymous/authenticate [get]
func (instance AuthHandlers) Login(context echo.Context) error {
	authURL, err := instance.service.Login()
	if err != nil {
		return getHttpHandledErrorResponse(context, err)
	}

	return context.JSON(http.StatusTemporaryRedirect, response.NewLoginRedirect(authURL))
}

// Callback
// @ID Callback
// @Summary Recebe a autorização da API do spotify
// @Tags Rotas anônimas
// @Description Rota que permite que um usuário se autentique no Echofy com seus dados de sua conta do Spotify.										  |
// @Accept json
// @Produce json
// @Param code path string true "Código de autorização"
// @Param state path string true "Valor de Estado requerido para segurança"
// @Success 200 {object} nil "Requisição Realizada com sucesso"
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /anonymous/callback [get]
func (instance AuthHandlers) Callback(context echo.Context) error {
	code := context.QueryParam("code")
	state := context.QueryParam("state")

	// TODO: Create invalidFields error messages in the messages file
	stateID, err := uuid.Parse(state)
	if err != nil {
		invalidField := errors.InvalidField{
			Name:        "O Código de Autorização é inválido!",
			Description: "Código de Autorização inválido! Adicione um código válido!",
		}
		validationErr := errors.NewValidationError("Código de Autorização", invalidField)
		return getHttpHandledErrorResponse(context, validationErr)
	}

	callbackErr := instance.service.Callback(code, stateID)
	if callbackErr != nil {
		return getHttpHandledErrorResponse(context, callbackErr)
	}

	homeURL := utils.GetenvWithDefault("HOME_URL", "http://localhost:5173/home")

	return context.Redirect(http.StatusPermanentRedirect, homeURL)
}

func NewAuthHandlers(service primary.AuthManager) *AuthHandlers {
	return &AuthHandlers{service: service}
}
