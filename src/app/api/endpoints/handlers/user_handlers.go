package handlers

import (
	"echofy_backend/src/app/api/endpoints/handlers/dtos/response"
	"echofy_backend/src/core/interfaces/primary"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	albumID    = "albumID"
	artistID   = "artistID"
	songID     = "songID"
	playlistID = "playlistID"
)

type UserHandlers struct {
	service primary.UserManager
}

// GetSongsByPlaylistID
// @ID GetSongsByPlaylistID
// @Summary Buscar todas as músicas de uma playlist
// @Tags Rotas do usuário
// @Description Rota que permite que se busque todas as músicas de uma determinada playlist
// @Param playlistID path string true "ID da playlist." default(7pCvSVfjcnOw6AFJNZZ4bN)
// @Produce json
// @Success 200 {array} response.SongDTO "Requisição realizada com sucesso."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /user/playlist/{playlistID}/songs [get]
func (h UserHandlers) GetSongsByPlaylistID(context echo.Context) error {
	playlistID := context.Param(playlistID)

	songRows, fetchErr := h.service.FetchSongsByPlaylistID(playlistID)
	if fetchErr != nil {
		return getHttpHandledErrorResponse(context, fetchErr)
	}

	songs := make([]response.SongDTO, 0)
	for _, each := range songRows {
		artists := make([]response.ArtistDTO, 0)
		for _, eachArtist := range each.Artists() {
			artistBuilder := response.NewArtistWithLowDataDTO(
				eachArtist.ID(),
				eachArtist.Name(),
				*eachArtist.SpotifyURL(),
			)

			artists = append(artists, *artistBuilder)
		}
		songBuilder := response.NewSongDTO(
			each.ID(),
			each.Name(),
			artists,
			each.AlbumID(),
			each.ReleaseDate(),
			each.Duration(),
		)
		songs = append(songs, *songBuilder)
	}

	return context.JSON(http.StatusOK, songs)
}

// GetPlaylistID
// @ID GetPlaylistID
// @Summary Buscar os dados de uma playlist
// @Tags Rotas do usuário
// @Description Rota que permite que se busque todas as informações de uma playlist
// @Param playlistID path string true "ID da playlist." default(7pCvSVfjcnOw6AFJNZZ4bN)
// @Produce json
// @Success 200 {object} response.PlaylistDTO "Requisição realizada com sucesso."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /user/playlist/{playlistID} [get]
func (h UserHandlers) GetPlaylistByID(context echo.Context) error {
	playlistID := context.Param(playlistID)

	playlistRow, fetchErr := h.service.FetchPlaylistByID(playlistID)
	if fetchErr != nil {
		return getHttpHandledErrorResponse(context, fetchErr)
	}

	playlist := response.NewPlaylistDTO(*playlistRow)

	return context.JSON(http.StatusOK, playlist)
}

// Authorize
// @ID Authorize
// @Summary Authorize
// @Tags Rotas anônimas
// @Description Authorize
// @Produce json
// @Success 200 {object} response.PlaylistDTO "Requisição realizada com sucesso."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /anonymous/authorize [get]
func (h UserHandlers) Authorize(context echo.Context) error {
    err := StartAuthorization(context)
    if err != nil {
        // Lide com erros, se necessário
        return context.JSON(http.StatusInternalServerError, "Erro ao iniciar o processo de autorização")
    }

    // O processo de autorização foi iniciado com sucesso, não é necessário retornar uma resposta aqui
    return nil
}


func NewUserHandlers(service primary.UserManager) *UserHandlers {
	return &UserHandlers{service: service}
}
