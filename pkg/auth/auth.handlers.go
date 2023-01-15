package pkg_auth

import (
	"errors"
	"fmt"
	"go-mono/configs"
	"go-mono/exception"
	"go-mono/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	Name, Password, Email string
}

type Result struct {
	Status  uint
	Message string
	Result  int64
}

func Register(c echo.Context) error {
	hashed, err := PasswordHashing(c.FormValue("password"))

	if err != nil {
		panic("Failed to hashing password")
	}

	user := model.User{
		Email:          c.FormValue("email"),
		Name:           c.FormValue("name"),
		HashedPassword: hashed,
	}
	result, err := CreateUser(&user)

	if err != nil {
		return exception.NotFound(err)
	}

	return c.JSON(http.StatusOK, &Result{
		Status:  http.StatusOK,
		Message: "Successfully created user " + user.Email,
		Result:  result.RowsAffected,
	})
}

func Login(c echo.Context) error {
	db := configs.ConnectDB("mysql")
	email := c.FormValue("email")
	password := c.FormValue("password")

	user := model.User{
		Email:          email,
		HashedPassword: password,
	}

	result := db.First(&user)

	errors.Is(result.Error, gorm.ErrRecordNotFound)

	fmt.Println(result)
	return c.JSON(http.StatusOK, &Result{
		Status:  http.StatusOK,
		Message: "Registered Users" + email,
		Result:  result.RowsAffected,
	})
}
