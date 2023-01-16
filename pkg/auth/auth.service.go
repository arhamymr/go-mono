package pkg_auth

import (
	"go-mono/configs"
	"go-mono/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type jwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (r ModelUser) CreateToken() (string, error) {
	claims := &jwtCustomClaims{
		r.data.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(configs.GET("SECRET")))

	if err != nil {
		return "", err
	}

	return t, nil
}

func PasswordHashing(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashed), err
}

func ComparePassword(hashed string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err
}

func (r ModelUser) CreateUser() (*model.User, error) {
	var db = configs.ConnectDB("mysql")
	result := db.Create(r.data)
	err := result.Error

	return r.data, err
}

func (r ModelUser) FindUser() (*model.User, error) {
	var db = configs.ConnectDB("mysql")
	result := db.Where("email =?", r.data.Email).First(&r.data)
	err := result.Error
	return r.data, err
}
