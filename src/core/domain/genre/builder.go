package genre

import (
	"echofy_backend/src/core/domain"
	"echofy_backend/src/core/errors"
	"echofy_backend/src/core/messages"
	"time"

	"github.com/google/uuid"
)

type Builder struct {
	Genre
	invalidFields []errors.InvalidField
}

func (b *Builder) WithID(id uuid.UUID) *Builder {
	if id == uuid.Nil {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.GenreID,
			Description: messages.GenreIDInvalidErrMsg,
		})
	}
	b.id = id
	return b
}

func (b *Builder) WithName(name string) *Builder {
	if name == "" {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.GenreName,
			Description: messages.GenreNameInvalidErrMsg,
		})
	}
	b.name = name
	return b
}

func (b *Builder) WithDescription(description *string) *Builder {
	b.description = description
	return b
}

func (b *Builder) WithCreatedAt(createdAt time.Time) *Builder {
	if createdAt.IsZero() {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.GenreCreatedAt,
			Description: messages.GenreCreatedAtInvalidErrMsg,
		})
	}
	b.createdAt = createdAt
	return b
}

func (b *Builder) Build() (*Genre, errors.Error) {
	if len(b.invalidFields) > 0 {
		domain.ShowInvalidFields(b.invalidFields)
		return nil, errors.NewValidationError(messages.GenreBuildErr, b.invalidFields...)
	}
	return &b.Genre, nil
}

func NewBuilder() *Builder {
	return &Builder{}
}
