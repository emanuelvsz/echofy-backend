package response

import (
	"time"

	"github.com/google/uuid"
)

type ArtistDTO struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	SuperArtistID *uuid.UUID  `json:"super_artist_id,omitempty"`
	Description   *string     `json:"description,omitempty"`
	FoundedAt     time.Time   `json:"founded_at,omitempty"`
	TerminatedAt  *time.Time  `json:"terminated_at,omitempty"`
	SubArtists    []ArtistDTO `json:"members,omitempty"`
	SpotifyURL    string      `json:"spotify_url,omitempty"`
}

func NewArtistDTO(id string, name string, superArtistID *uuid.UUID, description *string, foundedAt time.Time, terminatedAt *time.Time, subArtists []ArtistDTO) *ArtistDTO {
	return &ArtistDTO{
		ID:            id,
		Name:          name,
		SuperArtistID: superArtistID,
		Description:   description,
		FoundedAt:     foundedAt,
		TerminatedAt:  terminatedAt,
		SubArtists:    subArtists,
	}
}

type ArtistLowDTO struct {
	ID              uuid.UUID   `json:"id"`
	Name            string      `json:"name"`
	SuperArtistID   *uuid.UUID  `json:"super_artist_id,omitempty"`
	Description     *string     `json:"description,omitempty"`
	FoundedAt       time.Time   `json:"founded_at,omitempty"`
	TerminatedAt    *time.Time  `json:"terminated_at,omitempty"`
	SubArtists      []ArtistDTO `json:"members,omitempty"`
	ImageURL        *string     `json:"picture_url,omitempty"`
	RecordCompanyID *uuid.UUID  `json:"record_company_id,omitempty"`
	CountryID       *uuid.UUID  `json:"country_id,omitempty"`
	SpotifyURL      string      `json:"spotify_url,omitempty"`
}

func NewArtistLowDTO(id uuid.UUID, name string, superArtistID *uuid.UUID, description *string,
	foundedAt time.Time, terminatedAt *time.Time, subArtists []ArtistDTO, imageURL *string,
	recordCompanyID *uuid.UUID, countryID *uuid.UUID, spotifyURL string) *ArtistLowDTO {
	return &ArtistLowDTO{
		ID:              id,
		Name:            name,
		SuperArtistID:   superArtistID,
		Description:     description,
		FoundedAt:       foundedAt,
		TerminatedAt:    terminatedAt,
		SubArtists:      subArtists,
		ImageURL:        imageURL,
		RecordCompanyID: recordCompanyID,
		CountryID:       countryID,
		SpotifyURL:      spotifyURL,
	}
}

type ArtistWithLowDataDTO struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	SpotifyURL string `json:"spotify_url,omitempty"`
}

func NewArtistWithLowDataDTO(id, name, spotifyURL string) *ArtistDTO {
	return &ArtistDTO{
		ID:         id,
		Name:       name,
		SpotifyURL: spotifyURL,
	}
}
