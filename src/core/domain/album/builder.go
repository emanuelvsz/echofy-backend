package album

import (
	"echofy_backend/src/core/domain"
	"echofy_backend/src/core/errors"
	"echofy_backend/src/core/messages"
	"time"
)

type Builder struct {
	Album
	invalidFields []errors.InvalidField
}

func (b *Builder) WithID(id string) *Builder {
	if id == "" {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.AlbumID,
			Description: messages.AlbumIDInvalidErrMsg,
		})
	}
	b.id = id
	return b
}

func (b *Builder) WithName(name string) *Builder {
	if name == "" {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.AlbumName,
			Description: messages.AlbumNameInvalidErrMsg,
		})
	}
	b.name = name
	return b
}

func (b *Builder) WithArtistID(artistID string) *Builder {
	if artistID == "" {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.AlbumArtistID,
			Description: messages.AlbumArtistIDInvalidErrMsg,
		})
	}
	b.artistID = artistID
	return b
}

func (b *Builder) WithReleaseDate(releaseDate time.Time) *Builder {
	if releaseDate.IsZero() {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.AlbumReleaseDate,
			Description: messages.AlbumReleaseDateInvalidErrMsg,
		})
	}
	b.releaseDate = releaseDate
	return b
}

func (b *Builder) WithDescription(description string) *Builder {
	b.description = &description
	return b
}

// TODO: resolve this. resolve: put data in fixtures

func (b *Builder) WithImageURL(imageURL string) *Builder {
	// if imageURL == "" {
	// 	b.invalidFields = append(b.invalidFields, errors.InvalidField{
	// 		Name:        messages.AlbumImageURL,
	// 		Description: messages.AlbumImageURLInvalidErrMsg,
	// 	})
	// }
	b.imageURL = imageURL
	return b
}

func (b *Builder) Build() (*Album, errors.Error) {
	if len(b.invalidFields) > 0 {
		domain.ShowInvalidFields(b.invalidFields)
		return nil, errors.NewValidationError(messages.AlbumBuildErr, b.invalidFields...)
	}

	return &b.Album, nil
}

func NewBuilder() *Builder {
	return &Builder{}
}
