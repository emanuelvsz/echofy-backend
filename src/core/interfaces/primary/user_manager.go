package primary

import (
	"echofy_backend/src/core/domain/album"
	"echofy_backend/src/core/domain/playlist"
	"echofy_backend/src/core/domain/song"
	"echofy_backend/src/core/domain/user"
	"echofy_backend/src/core/errors"
)

type UserManager interface {
	FetchSongsByPlaylistID(playlistID string) ([]song.Song, errors.Error)
	FetchPlaylistByID(playlistID string) (*playlist.Playlist, errors.Error)
	FetchSongsByAlbumID(albumID string) ([]song.Song, errors.Error)
	FetchUserBasicInfo() (*user.User, errors.Error)
	FetchArtistAlbumsByID(artistID string) ([]album.Album, errors.Error)
}
