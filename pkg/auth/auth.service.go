package pkg_auth

import (
	"go-mono/configs"
	"go-mono/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func PasswordHashing(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashed), err
}

func CreateUser(data *model.User) (*gorm.DB, error) {
	var db = configs.ConnectDB("mysql")
	result := db.Create(&data)
	err := result.Error

	return result, err
}
