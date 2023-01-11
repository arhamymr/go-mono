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

func PasswordHashing(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashed), err
}

func Register(c echo.Context) error {

	db := configs.ConnectDB("mysql")

	hashed, err := PasswordHashing(c.FormValue("password"))

	if err != nil {
		panic("Failed to hashing password")
	}

	u := model.User{
		Email:          c.FormValue("email"),
		Name:           c.FormValue("name"),
		HashedPassword: hashed,
	}

	err = db.Create(&u).Error

	if err != nil {
		type ErrorMessage struct {
			Status  uint
			Message string
		}

		return echo.NewHTTPError(http.StatusBadRequest, ErrorMessage{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	type Result struct {
		Status  uint
		Message string
	}

	return c.JSON(200, Result{
		Status:  http.StatusOK,
		Message: "Successfully created user" + u.Email,
	})

}

func Login(c echo.Context) error {
	return c.String(http.StatusOK, "hello login")
}
