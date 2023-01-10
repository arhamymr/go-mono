package main

import (
	users "go-mono/handler"
	"net/http"

	"github.com/labstack/echo"
)

const PORT = ":3000"

func getUser(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func main() {
	e := echo.New()
	e.Static("/static", "static")
	e.GET("/", getUser)

	auth := e.Group("/auth")
	auth.POST("/register", users.Register)
	auth.POST("/login", users.Login)

	e.Logger.Fatal(e.Start(PORT))

}
