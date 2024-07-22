package song

import (
	"echofy_backend/src/core/domain/artist"
	"time"
)

type Song struct {
	id          string
	name        string
	artists     []artist.Artist
	albumID     string
	releaseDate time.Time
	duration    int
	lyrics      *string
	trackNumber *int
	spotifyURL  string
}

func (s Song) ID() string {
	return s.id
}

func (s Song) Name() string {
	return s.name
}

func (s Song) Artists() []artist.Artist {
	return s.artists
}

func (s Song) AlbumID() string {
	return s.albumID
}

func (s Song) ReleaseDate() time.Time {
	return s.releaseDate
}

func (s Song) Duration() int {
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
