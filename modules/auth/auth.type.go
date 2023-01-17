package mod_auth

import (
	"go-mono/model"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Name, Password, Email string
}

type Result struct {
	Status  uint
	Message string
	Result  *model.User
	Token   string
}

type ModelUser struct {
	data *model.User
}

type jwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
