package pkg_auth

import (
	"go-mono/exception"
	"go-mono/model"
	"net/http"

	"github.com/labstack/echo/v4"
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

	data := ModelUser{&user}
	result, err := data.CreateUser()

	if err != nil {
		return exception.NotFound(err)
	}

	return c.JSON(http.StatusOK, &Result{
		Status:  http.StatusOK,
		Message: "Successfully created user " + user.Email,
		Result:  result,
	})
}

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	hashed, err := PasswordHashing(password)
	user := model.User{
		Email:          email,
		HashedPassword: hashed,
	}

	// find user
	data := ModelUser{&user}
	result, err := data.FindUser()

	if err != nil {
		return exception.NotFound(err)
	}

	// Compare password
	err = ComparePassword(result.HashedPassword, password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "password not found")
	}

	// Generate token
	token, err := data.CreateToken()
	if err != nil {
		panic("failed to create token")
	}

	return c.JSON(http.StatusOK, &Result{
		Status:  http.StatusOK,
		Message: "Registered User " + email,
		Token:   token,
		Result:  result,
	})
}
