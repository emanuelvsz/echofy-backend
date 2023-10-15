package repository

import (
	"echofy_backend/src/core/domain/playlist"
	"echofy_backend/src/core/domain/song"
	"echofy_backend/src/core/domain/user"
	"echofy_backend/src/core/errors"
)

type UserLoader interface {
	FindSongsByPlaylistID(playlistID string) ([]song.Song, errors.Error)
	FindPlaylistByID(playlistID string) (*playlist.Playlist, errors.Error)
	FindSongsByAlbumID(albumID string) ([]song.Song, errors.Error)
	FindUserBasicInfo() (*user.User, errors.Error)
}
