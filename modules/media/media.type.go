package mod_media

import (
	"go-mono/model"

	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type Result struct {
	Status  uint
	Message string
	Result  *model.Post
}
