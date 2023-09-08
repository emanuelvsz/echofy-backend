package state

import "github.com/google/uuid"

func GenerateRandomState() string {
	uuidObj := uuid.New()
	return uuidObj.String()
}
