package response

import "echofy_backend/src/core/domain/user"

type UserDTO struct {
	ID         *string `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	PictureURL string  `json:"picture_url"`
	URI        string  `json:"uri"`
}

func NewUserDTO(user user.User) *UserDTO {
	return &UserDTO{
		ID:         user.ID(),
		Name:       user.Name(),
		Email:      user.Email(),
		PictureURL: user.PictureURL(),
		URI:        user.URI(),
	}
}
