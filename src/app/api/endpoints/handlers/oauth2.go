package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

var oauth2Config = oauth2.Config{
	ClientID:     "3348f4b437614e8b9c742c305eb9865b",
	ClientSecret: "cd3d51d52a724d18aac6d3910534420c",
	RedirectURL:  "http://localhost:8000/api/",
	Scopes:       []string{"user-read-private", "user-read-email"},
	Endpoint:     spotify.Endpoint,
}

// Função para gerar um estado aleatório
func generateRandomState() string {
	const stateSize = 32
	b := make([]byte, stateSize)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("Erro ao gerar estado aleatório: %v", err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

func StartAuthorization(context echo.Context) error {
	// Crie um estado aleatório para evitar ataques CSRF
	state := generateRandomState()

	// Armazene o estado em algum lugar seguro, como no banco de dados ou na sessão do usuário
	// Isso será usado para verificar a correspondência durante o callback

	// Redirecione o usuário para a página de autorização do Spotify
	authorizationURL := oauth2Config.AuthCodeURL(state)
	return context.Redirect(http.StatusFound, authorizationURL)
}

func SpotifyCallback(context echo.Context) error {
	code := context.QueryParam("code")   // Obtém o código de autorização do Spotify
	state := context.QueryParam("state") // Obtém o estado retornado pelo Spotify

	fmt.Println(state)
	// Use o código de autorização para obter o Token de Acesso
	token, err := oauth2Config.Exchange(context.Request().Context(), code)
	if err != nil {
		log.Println("Erro ao obter o Token de Acesso:", err)
		return context.JSON(http.StatusInternalServerError, "Erro ao obter o Token de Acesso")
	}

	// Agora, você tem o Token de Acesso e pode usá-lo para acessar os dados do usuário
	// Você pode armazenar o Token de Acesso em algum lugar seguro para usos futuros, como em um banco de dados

	// Redirecione ou retorne uma resposta, dependendo do que deseja fazer em seguida
	return context.JSON(http.StatusOK, token)
}
