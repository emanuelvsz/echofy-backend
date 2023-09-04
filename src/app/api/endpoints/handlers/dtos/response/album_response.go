package response

import (
	"time"

	"github.com/google/uuid"
)

type AlbumDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ArtistID    uuid.UUID `json:"artistId"`
	ReleaseDate time.Time `json:"release_date"`
	Description *string   `json:"description,omitempty"`
	ImageURL    *string   `json:"image_url,omitempty"`
}

func NewAlbumDTO(id uuid.UUID, name string, artistID uuid.UUID, releaseDate time.Time, description *string, imageURL *string) *AlbumDTO {
	return &AlbumDTO{
		ID:          id,
		Name:        name,
		ArtistID:    artistID,
		ReleaseDate: releaseDate,
		Description: description,
		ImageURL:    imageURL,
	}
}
