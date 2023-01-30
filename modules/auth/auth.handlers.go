package mod_auth

import (
	"go-mono/exception"
	"go-mono/jwt"
	"go-mono/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

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
		return exception.CustomException(err)
	}

	return c.JSON(http.StatusOK, &Result{
		Status:  http.StatusOK,
		Message: "Successfully created user " + user.Email,
		Result:  result,
	})
}

func Login(c echo.Context) (err error) {
	email := c.FormValue("email")
	password := c.FormValue("password")

	dto := &LoginValidation{
		Email:    email,
		Password: password,
	}

	if err = c.Bind(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(dto); err != nil {
		return err
	}

	hashed, err := PasswordHashing(dto.Password)
	user := model.User{
		Email:          dto.Email,
		HashedPassword: hashed,
	}

	// find user
	data := ModelUser{&user}
	result, err := data.FindUser()

	if err != nil {
		return exception.CustomException(err)
	}

	// Compare password
	err = ComparePassword(result.HashedPassword, password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "password not found")
	}

	// Generate token
	token, err := jwt.CreateToken(&user)

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, &Result{
		Status:  http.StatusOK,
		Message: "Registered User " + email,
		Token:   token,
		Result:  result,
	})
}
