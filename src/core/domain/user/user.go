package user

type User struct {
	id         *string
	name       string
	pictureURL string
	email      string
	uri        string
}

func (u User) ID() *string {
	return u.id
}

func (u User) Name() string {
	return u.name
}

func (u User) PictureURL() string {
	return u.pictureURL
}

func (u User) Email() string {
	return u.email
}

func (u User) URI() string {
	return u.uri
}
