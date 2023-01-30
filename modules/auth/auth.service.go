package mod_auth

import (
	"go-mono/configs"
	"go-mono/model"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHashing(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashed), err
}

func ComparePassword(hashed string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err
}

func (r ModelUser) CreateUser() (*model.User, error) {
	result := configs.DB.Create(r.data)
	err := result.Error

	return r.data, err
}

func (r ModelUser) FindUser() (*model.User, error) {
	result := configs.DB.Where("email =?", r.data.Email).First(&r.data)
	err := result.Error
	return r.data, err
}
