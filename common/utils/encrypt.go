package utils

import (
	"uminer/common/errors"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Errorf(err, errors.ErrorEncryptPasswordFailed)
	}
	return string(pwd), nil
}

func ValidatePassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
