package user

import (
	"echofy_backend/src/core/domain"
	"echofy_backend/src/core/errors"
	"echofy_backend/src/core/messages"
	"strings"
)

type Builder struct {
	User
	invalidFields []errors.InvalidField
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) WithID(id *string) *Builder {
	b.id = id
	return b
}

func (b *Builder) WithName(name string) *Builder {
	name = strings.TrimSpace(name)
	if name == "" {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.UserName,
			Description: messages.UserNameInvalidErrMsg,
		})
	}
	b.name = name
	return b
}

func (b *Builder) WithPictureURL(pictureURL string) *Builder {
	b.pictureURL = pictureURL
	return b
}

func (b *Builder) WithEmail(email string) *Builder {
	b.email = email
	return b
}

func (b *Builder) WithURI(uri string) *Builder {
	b.uri = uri
	return b
}

func (b *Builder) Build() (*User, errors.Error) {
	if len(b.invalidFields) > 0 {
		domain.ShowInvalidFields(b.invalidFields)
		return nil, errors.NewValidationError(messages.UserBuildErr, b.invalidFields...)
	}

	return &User{
		id:         b.id,
		name:       b.name,
		pictureURL: b.pictureURL,
		email:      b.email,
		uri:        b.uri,
	}, nil
}
