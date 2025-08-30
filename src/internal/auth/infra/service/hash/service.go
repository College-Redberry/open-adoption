package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type HashService struct{}

func New() *HashService {
	return &HashService{}
}

func (service *HashService) Hash(text string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func (service *HashService) Verify(text string, hashedText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(text))
	if err != nil {
		return false, err
	}

	return true, nil
}
