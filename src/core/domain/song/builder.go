package song

import (
	"echofy_backend/src/core/domain"
	"echofy_backend/src/core/domain/artist"
	"echofy_backend/src/core/errors"
	"echofy_backend/src/core/messages"
	"net/url"
	"time"

	"github.com/google/uuid"
)

type Builder struct {
	Song
	invalidFields []errors.InvalidField
}

func (b *Builder) WithID(id string) *Builder {
	if id == "" {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.SongID,
			Description: messages.SongIDInvalidErrMsg,
		})
	} else {
		b.id = id
	}
	return b
}

func (b *Builder) WithName(name string) *Builder {
	if name == "" {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.SongName,
			Description: messages.SongNameInvalidErrMsg,
		})
	} else {
		b.name = name
	}
	return b
}

func (b *Builder) WithArtists(artists []artist.Artist) *Builder {
	if artists == nil {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.SongArtistID,
			Description: messages.SongArtistIDInvalidErrMsg,
		})
	} else {
		b.artists = artists
	}
	return b
}

func (b *Builder) WithAlbumID(albumID *uuid.UUID) *Builder {
	b.albumID = albumID
	return b
}

func (b *Builder) WithReleaseDate(releaseDate time.Time) *Builder {
	if releaseDate.IsZero() {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.SongReleaseDate,
			Description: messages.SongReleaseDateInvalidErrMsg,
		})
	} else {
		b.releaseDate = releaseDate
	}
	return b
}

func (b *Builder) WithDuration(duration int) *Builder {
	if duration == 0 {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.SongDuration,
			Description: messages.SongDurationInvalidErrMsg,
		})
	} else {
		b.duration = duration
	}
	return b
}

func (b *Builder) WithLyrics(lyrics string) *Builder {
	b.lyrics = &lyrics
	return b
}

// TODO: resolve this two functions

func (b *Builder) WithTrackNumber(trackNumber int) *Builder {
	// if trackNumber <= 0 {
	// 	b.invalidFields = append(b.invalidFields, errors.InvalidField{
	// 		Name:        messages.SongTrackNumber,
	// 		Description: messages.SongTrackNumberInvalidErrMsg,
	// 	})
	// }
	b.trackNumber = &trackNumber
	return b
}

func (b *Builder) WithSpotifyURL(spotifyURL string) *Builder {
	// if !IsValidURL(spotifyURL) {
	// 	b.invalidFields = append(b.invalidFields, errors.InvalidField{
	// 		Name:        messages.SongSpotifyURL,
	// 		Description: messages.SongSpotifyURLInvalidErrMsg,
	// 	})
	// }
	b.spotifyURL = spotifyURL
	return b
}

func (b *Builder) Build() (*Song, errors.Error) {
	if len(b.invalidFields) > 0 {
		domain.ShowInvalidFields(b.invalidFields)
		return nil, errors.NewValidationError(messages.SongBuildErr, b.invalidFields...)
	}

	return &b.Song, nil
}

func IsValidURL(str string) bool {
	_, err := url.Parse(str)
	return err != nil
}

func NewBuilder() *Builder {
	return &Builder{}
}
