package pkg_posts

import (
	"fmt"
	"go-mono/exception"
	"go-mono/jwt"
	"go-mono/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	token := jwt.ExtractToken(c)

	fmt.Println(token, "token")
	post := model.Post{
		Title:     c.FormValue("title"),
		Content:   c.FormValue("content"),
		Published: true,
		AuthorID:  token.ID,
	}

	data := ModelPost{data: &post}
	err := data.CreatePost()

	if err != nil {
		return exception.CustomException(err)
	}

	return c.JSON(200, Result{
		Status:  http.StatusOK,
		Message: "Posts Created",
		Result:  &post,
	})
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
