package encrypt

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, string, error) {
	salt := generateSalt()
	bytes := applySaltToPassword(password, salt)

	encryptedPassword, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("There was an error encrypting the password: %v", err.Error())
		return "", "", err
	}

	return encodeHashToString(encryptedPassword), encodeHashToString(salt), nil
}

func PasswordsMatch(password, hash, encryptedPassword string) bool {
	salt, err := decodeHashFromString(hash)
	if err != nil {
		return false
	}

	hashedPassword, err := decodeHashFromString(encryptedPassword)
	if err != nil {
		return false
	}

	bytes := applySaltToPassword(password, salt)
	err = bcrypt.CompareHashAndPassword(hashedPassword, bytes)
	return err == nil
}

func generateSalt() []byte {
	size := 256
	bytes := make([]byte, size)

	_, err := rand.Read(bytes)
	if err != nil {
		log.Errorf("There was an error generating the salt: %v", err.Error())
	}

	return bytes
}

func applySaltToPassword(password string, salt []byte) []byte {
	return append([]byte(password), salt[:]...)
}

func encodeHashToString(hash []byte) string {
	return hex.EncodeToString(hash)
}

func decodeHashFromString(hash string) ([]byte, error) {
	decoded, err := hex.DecodeString(hash)
	if err != nil {
		log.Errorf("There was an error decoding the password: %v", err.Error())
		return nil, err
	}

	return decoded, nil
}
