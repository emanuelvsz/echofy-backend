package user

import (
	"github.com/google/uuid"
)

type User struct {
	id         uuid.UUID
	name       string
	pictureURL string
	password   string
	hash       string
	email      string
}

func (u User) ID() uuid.UUID {
	return u.id
}

func (u User) Name() string {
	return u.name
}

func (u User) Hash() string {
	return u.hash
}

func (u User) Password() string {
	return u.password
}

func (u User) PictureURL() string {
	return u.pictureURL
}

func (u User) Email() string {
	return u.email
}