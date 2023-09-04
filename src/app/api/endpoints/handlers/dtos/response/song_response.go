package response

import (
	"time"

	"github.com/google/uuid"
)

type SongDTO struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Artists     []ArtistDTO `json:"artists"`
	AlbumID     *uuid.UUID  `json:"album_id,omitempty"`
	ReleaseDate time.Time   `json:"release_date,omitempty"`
	Duration    int         `json:"duration"`
}

func NewSongDTO(ID, Name string, Artists []ArtistDTO, AlbumID *uuid.UUID, ReleaseDate time.Time, Duration int) *SongDTO {
	return &SongDTO{
		ID:          ID,
		Name:        Name,
		Artists:     Artists,
		AlbumID:     AlbumID,
		ReleaseDate: ReleaseDate,
		Duration:    Duration,
	}
}

type SongHighDTO struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	AlbumID     *uuid.UUID `json:"album_id,omitempty"`
	ReleaseDate time.Time  `json:"release_date"`
	Duration    string     `json:"duration"`
	Lyrics      *string    `json:"lyrics"`
	TrackNumber *int       `json:"track_number"`
	SpotifyURL  *string    `json:"spotify_url"`
}

func NewSongHighDTO(ID uuid.UUID, Name string, AlbumID *uuid.UUID,
	ReleaseDate time.Time, Duration string, Lyrics, SpotifyURL *string, TrackNumber *int) *SongHighDTO {
	return &SongHighDTO{
		ID:          ID,
		Name:        Name,
		AlbumID:     AlbumID,
		ReleaseDate: ReleaseDate,
		Duration:    Duration,
		Lyrics:      Lyrics,
		TrackNumber: TrackNumber,
		SpotifyURL:  SpotifyURL,
	}
}
