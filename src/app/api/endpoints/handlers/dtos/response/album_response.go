package response

import (
	"time"
)

type AlbumDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	ArtistID    string    `json:"artistId"`
	ReleaseDate time.Time `json:"release_date"`
	Description *string   `json:"description,omitempty"`
	ImageURL    *string   `json:"image_url,omitempty"`
}

func NewAlbumDTO(id string, name string, artistID string, releaseDate time.Time, description *string, imageURL *string) *AlbumDTO {
	return &AlbumDTO{
		ID:          id,
		Name:        name,
		ArtistID:    artistID,
		ReleaseDate: releaseDate,
		Description: description,
		ImageURL:    imageURL,
	}
}
