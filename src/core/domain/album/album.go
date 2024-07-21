package album

import (
	"time"
)

type Album struct {
	id          string
	name        string
	artistID    string
	releaseDate time.Time
	description *string
	imageURL    string
}

func (a Album) ID() string {
	return a.id
}

func (a Album) Name() string {
	return a.name
}

func (a Album) ArtistID() string {
	return a.artistID
}

func (a Album) ReleaseDate() time.Time {
	return a.releaseDate
}

func (a Album) Description() *string {
	return a.description
}

func (a Album) ImageURL() *string {
	return &a.imageURL
}
