package handlers

import (
	"echofy_backend/src/app/api/endpoints/handlers/dtos/response"
	"echofy_backend/src/core/interfaces/primary"
	"net/http"
	"time"

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

// GetAlbumTracks
// @ID GetAlbumTracks
// @Summary Buscar todas as músicas de um Álbum
// @Tags Rotas do usuário
// @Description Rota que permite que se busque todas as músicas de uma determinado Álbum
// @Param albumID path string true "ID do Álbum." default(3WFTGIO6E3Xh4paEOBY9OU)
// @Produce json
// @Success 200 {array} response.SongDTO "Requisição realizada com sucesso."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /user/album/{albumID}/songs [get]
func (h UserHandlers) GetAlbumTracks(context echo.Context) error {
	albumID := context.Param(albumID)

	songRows, fetchErr := h.service.FetchSongsByAlbumID(albumID)
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

// GetAlbumByArtistID
// @ID GetAlbumByArtistID
// @Summary Buscar todas os álbuns de um artista pelo seu ID
// @Tags Rotas do usuário
// @Description Rota que permite que se busque todos os álbuns de um determinado artista
// @Param artistID path string true "ID do Artista." default(5K4W6rqBFWDnAN6FQUkS6x)
// @Produce json
// @Success 200 {array} response.AlbumDTO "Requisição realizada com sucesso."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /user/artist/{artistID}/albums [get]
func (h UserHandlers) GetArtistAlbums(context echo.Context) error {
	artistID := context.Param(artistID)

	albumsRows, fetchErr := h.service.FetchArtistAlbumsByID(artistID)
	if fetchErr != nil {
		return getHttpHandledErrorResponse(context, fetchErr)
	}

	albums := make([]response.AlbumDTO, 0)
	for _, each := range albumsRows {
		albumBuilder := response.NewAlbumDTO(
			each.ID(),
			each.Name(),
			each.ArtistID(),
			each.ReleaseDate(),
			each.Description(),
			each.ImageURL(),
		)
		albums = append(albums, *albumBuilder)

	}

	return context.JSON(http.StatusOK, albums)

}

// GetSongDetails
// @ID GetSongDetails
// @Summary Buscar dados de uma música especifíca
// @Tags Rotas do usuário
// @Description Rota que permite que se busque os dados de uma música
// @Param songID path string true "ID da música." default(1SKPmfSYaPsETbRHaiA18G)
// @Produce json
// @Success 200 {array} response.SongDTO "Requisição realizada com sucesso."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /user/song/{songID}/details [get]
func (h UserHandlers) FetchSongDetailsByID(context echo.Context) error {
	songID := context.Param("songID")

	songRows, fetchErr := h.service.FetchSongDetailsByID(songID)
	if fetchErr != nil {
		return getHttpHandledErrorResponse(context, fetchErr)
	}

	artists := make([]response.ArtistDTO, 0)
	for _, each := range songRows.Artists() {
		artistBuilder := response.NewArtistDTO(
			each.ID(),
			each.Name(),
			nil,
			nil,
			time.Time{},
			nil,
			nil,
		)
		artists = append(artists, *artistBuilder)
	}

	track := response.NewSongDTO(
		songRows.ID(),
		songRows.Name(),
		artists,
		songRows.AlbumID(),
		songRows.ReleaseDate(),
		songRows.Duration(),
	)

	return context.JSON(http.StatusOK, track)

}

// GetUserBasicInfo
// @ID GetUserBasicInfo
// @Summary Buscar alguns dados pessoais do usuario
// @Tags Rotas do usuário
// @Description Rota que permite que se busque alguns dados do usuário autenticado
// @Produce json
// @Success 200 {array} response.UserDTO "Requisição realizada com sucesso."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /user [get]
func (h UserHandlers) GetUserBasicInfo(context echo.Context) error {
	userRow, fetchErr := h.service.FetchUserBasicInfo()
	if fetchErr != nil {
		return getHttpHandledErrorResponse(context, fetchErr)
	}
	user := response.NewUserDTO(*userRow)
	return context.JSON(http.StatusOK, user)
}

func NewUserHandlers(service primary.UserManager) *UserHandlers {
	return &UserHandlers{service: service}
}
