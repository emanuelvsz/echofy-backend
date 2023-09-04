package repository

import (
	"echofy_backend/src/core/domain/song"
	"echofy_backend/src/core/errors"
)

type UserLoader interface {
	FindSongsByPlaylistID(playlistID string) ([]song.Song, errors.Error)
}
