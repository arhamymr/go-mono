package jwt

import (
	"go-mono/configs"
	"go-mono/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func CreateToken(data *model.User) (string, error) {
	claims := &JwtCustomClaims{
		data.ID,
		data.Email,
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

func ExtractToken(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}
