package pkg_posts

import (
	"fmt"
	"go-mono/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	u := model.User{
		Email:          "arhamymr@gmail.com",
		Name:           "arham",
		HashedPassword: "hellow youth",
	}

	fmt.Println(u)
	return c.JSON(200, "test")
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
