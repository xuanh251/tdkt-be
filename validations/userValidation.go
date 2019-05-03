package validations

import (
	"TDKTServer/models"
	"errors"
)

var (
	ErrEmptyFields  = errors.New("One or more fields is empty")
	ErrInvalidEmail = errors.New("Invalid email format")
)

func ValidationNewUser(account models.Account) (models.Account, error) {
	if IsEmpty(account.UserName) || IsEmpty(account.Password) || IsEmpty(account.Email) {
		return models.Account{}, ErrEmptyFields
	}
	if !IsEmail(account.Email) {
		return models.Account{}, ErrInvalidEmail
	}
	return account, nil
}
