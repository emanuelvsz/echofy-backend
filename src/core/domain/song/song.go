package song

import (
	"time"

	"github.com/google/uuid"
)

type Song struct {
	id          uuid.UUID
	name        string
	artistID    uuid.UUID
	albumID     *uuid.UUID
	releaseDate time.Time
	duration    string
	lyrics      *string
	trackNumber *int
	spotifyURL  string
}


func (s Song) ID() uuid.UUID {
	return s.id
}

func (s Song) Name() string {
	return s.name
}

func (s Song) ArtistID() uuid.UUID {
	return s.artistID
}

func (s Song) AlbumID() *uuid.UUID {
	return s.albumID
}

func (s Song) ReleaseDate() time.Time {
	return s.releaseDate
}

func (s Song) Duration() string {
	return s.duration
}

func (s Song) Lyrics() *string {
	return s.lyrics
}

func (s Song) TrackNumber() *int {
	return s.trackNumber
}

func (s Song) SpotifyURL() string {
	return s.spotifyURL
}
