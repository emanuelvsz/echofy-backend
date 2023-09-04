package credentials

import "echofy_backend/src/core/errors"

type Credentials struct {
	email    string
	password string
}

func (instance Credentials) Email() string {
	return instance.email
}

func (instance Credentials) Password() string {
	return instance.password
}

func New(email string, password string) (*Credentials, errors.Error) {
	// TODO: validate email (by Leonardo)
	// TODO: validate password (by Leonardo)

	return &Credentials{
		email:    email,
		password: password,
	}, nil
}
