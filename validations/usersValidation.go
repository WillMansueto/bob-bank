package validations

import (
	"errors"

	"bob-bank/models"
)

var (
	ErrEmptyFields = errors.New("Um ou mais campos vazios")
	ErrInvalidEmail = errors.New("Email inválido")
)

func ValidateNewUser(user models.User) (models.User, error) {
	if IsEmpty(user.Nickname) || IsEmpty(user.Email) || IsEmpty(user.Password) {
		return models.User{}, ErrEmptyFields
	}
	if !IsMail(user.Email) {
		return models.User{}, ErrInvalidEmail
	}
	return user, nil
}