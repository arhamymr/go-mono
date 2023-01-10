package main

import (
	"net/http"

	"github.com/labstack/echo"
)

const PORT = ":3000"

func getUser(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func main() {
	e := echo.New()
	e.GET("/", getUser)
	e.Logger.Fatal(e.Start(PORT))
}
