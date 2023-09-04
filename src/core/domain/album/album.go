package album

import (
	"time"

	"github.com/google/uuid"
)

type Album struct {
	id          uuid.UUID
	name        string
	artistID    uuid.UUID
	releaseDate time.Time
	description *string
	imageURL    string
}

func (a Album) ID() uuid.UUID {
	return a.id
}

func (a Album) Name() string {
	return a.name
}

func (a Album) ArtistID() uuid.UUID {
	return a.artistID
}

func (a Album) ReleaseDate() time.Time {
	return a.releaseDate
}

func (a Album) Description() *string {
	return a.description
}

func (a Album) ImageURL() string {
	return a.imageURL
}
