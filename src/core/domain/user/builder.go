package user

import (
	"echofy_backend/src/core/domain"
	"echofy_backend/src/core/errors"
	"echofy_backend/src/core/messages"
	"strings"

	"github.com/google/uuid"
)

type Builder struct {
	User
	invalidFields []errors.InvalidField
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) WithID(id uuid.UUID) *Builder {

	if id == uuid.Nil {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.UserID,
			Description: messages.UserIDInvalidErrMsg,
		})
	}

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

	pictureURL = strings.TrimSpace(pictureURL)

	b.pictureURL = pictureURL
	return b
}

func (b *Builder) WithPassword(password string) *Builder {

	password = strings.TrimSpace(password)

	if password == "" {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.UserPassword,
			Description: messages.UserPasswordInvalidErrMsg,
		})
	}

	b.password = password
	return b
}

func (b *Builder) WithHash(hash string) *Builder {

	hash = strings.TrimSpace(hash)

	if hash == "" {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.UserName,
			Description: messages.UserNameInvalidErrMsg,
		})
	}

	b.hash = hash
	return b
}

func (b *Builder) WithEmail(email string) *Builder {

	email = strings.TrimSpace(email)

	if email == "" {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.UserEmail,
			Description: messages.UserEmailInvalidErrMsg,
		})
	}

	b.email = email
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
		password:   b.password,
		hash:       b.hash,
		email:      b.email,
	}, nil
}
