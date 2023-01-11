package handlers

import (
	"go-mono/configs"
	"go-mono/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name, Password, Email string
}

func PasswordHashing(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic("Failed to generate hashed password ")
	}
	return string(hashed)
}

func Register(c echo.Context) error {
	db := configs.ConnectDB("mysql")

	u := model.User{
		Email:          c.FormValue("email"),
		Name:           c.FormValue("name"),
		HashedPassword: PasswordHashing(c.FormValue("password")),
	}

	err := db.Create(&u)

	if err != nil {
		type ErrorMessage struct {
			Status  uint
			Message string
		}

		return echo.NewHTTPError(http.StatusBadRequest, ErrorMessage{
			Status:  http.StatusBadRequest,
			Message: err.Error.Error(),
		})
	}

	return c.JSON(200, u)

}

func Login(c echo.Context) error {
	return c.String(http.StatusOK, "hello login")
}
