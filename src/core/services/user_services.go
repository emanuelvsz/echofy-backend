package services

import (
	"echofy_backend/src/core/domain/playlist"
	"echofy_backend/src/core/domain/song"
	"echofy_backend/src/core/domain/user"
	"echofy_backend/src/core/errors"
	"echofy_backend/src/core/errors/logger"
	"echofy_backend/src/core/interfaces/primary"
	"echofy_backend/src/core/interfaces/repository"
)

var _ primary.UserManager = (*UserServices)(nil)

type UserServices struct {
	userRepository repository.UserLoader
	logger         logger.Logger
}

func (u UserServices) FetchSongsByPlaylistID(playlistID string) ([]song.Song, errors.Error) {
	songs, err := u.userRepository.FindSongsByPlaylistID(playlistID)
	if err != nil {
		return nil, err
	}

	return songs, nil
}

func (u UserServices) FetchPlaylistByID(playlistID string) (*playlist.Playlist, errors.Error) {
	playlist, err := u.userRepository.FindPlaylistByID(playlistID)
	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func (u UserServices) FetchSongsByAlbumID(albumID string) ([]song.Song, errors.Error) {
	album, err := u.userRepository.FindSongsByAlbumID(albumID)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (u UserServices) FetchUserBasicInfo() (*user.User, errors.Error) {
	user, err := u.userRepository.FindUserBasicInfo()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserServices(userRepository repository.UserLoader, logger logger.Logger) *UserServices {
	return &UserServices{
		userRepository: userRepository,
		logger:         logger,
	}
}
