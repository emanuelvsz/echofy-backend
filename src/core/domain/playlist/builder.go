package playlist

import (
	"echofy_backend/src/core/domain"
	"echofy_backend/src/core/errors"
	"echofy_backend/src/core/messages"
)

type Builder struct {
	Playlist
	invalidFields []errors.InvalidField
}

func (instance *Builder) WithID(id string) *Builder {
	if id == "" {
		instance.invalidFields = append(instance.invalidFields, errors.InvalidField{
			Name:        messages.PlaylistID,
			Description: messages.PlaylistIDInvalidErrMsg,
		})
	} else {
		instance.id = id
	}

	return instance
}

func (instance *Builder) WithName(name string) *Builder {
	if name == "" {
		instance.invalidFields = append(instance.invalidFields, errors.InvalidField{
			Name:        messages.PlaylistName,
			Description: messages.PlaylistNameInvalidErrMsg,
		})
	}
	instance.name = name

	return instance
}

func (instance *Builder) WithDescription(desc string) *Builder {
	instance.description = desc

	return instance
}

func (instance *Builder) WithSongAmount(amount int) *Builder {
	instance.songAmount = amount

	return instance
}

func (instance *Builder) WithFollowersAmount(amount int) *Builder {
	instance.followersAmount = amount

	return instance
}

func (b *Builder) Build() (*Playlist, errors.Error) {
	if len(b.invalidFields) > 0 {
		domain.ShowInvalidFields(b.invalidFields)
		return nil, errors.NewValidationError(messages.GenreBuildErr, b.invalidFields...)
	}
	return &b.Playlist, nil
}

func NewBuilder() *Builder {
	return &Builder{}
}
