package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Update(c echo.Context) error {
	u := User{"arham", "hashedpassword", "email"}
	return c.JSON(200, u)
}

func GetMany(c echo.Context) error {
	return c.String(http.StatusOK, "hello login")
}

func GetOne(c echo.Context) error {
	return c.String(http.StatusOK, "hello login")
}

func Delete(c echo.Context) error {
	return c.String(http.StatusOK, "hello login")
}
