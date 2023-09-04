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
// @Param playlistID path string true "ID do artista." default(7pCvSVfjcnOw6AFJNZZ4bN)
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

func NewUserHandlers(service primary.UserManager) *UserHandlers {
	return &UserHandlers{service: service}
}
