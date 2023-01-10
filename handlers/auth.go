package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name, Password, Email string
}

func Register(c echo.Context) error {
	u := User{"arham", "hashedpassword", "email"}
	return c.JSON(200, u)
}

func Login(c echo.Context) error {
	return c.String(http.StatusOK, "hello login")
}
