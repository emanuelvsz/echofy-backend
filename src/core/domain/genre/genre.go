package genre

import (
	"time"

	"github.com/google/uuid"
)

type Genre struct {
	id          uuid.UUID
	name        string
	description *string
	createdAt   time.Time
}

func (g Genre) ID() uuid.UUID {
	return g.id
}

func (g Genre) Name() string {
	return g.name
}

func (g Genre) Description() *string {
	return g.description
}

func (g Genre) CreatedAt() time.Time {
	return g.createdAt
}
