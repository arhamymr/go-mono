package jwt

import "github.com/golang-jwt/jwt/v4"

type JwtCustomClaims struct {
	ID    int32  `json:"ID"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
