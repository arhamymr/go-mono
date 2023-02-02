package main

import (
	"go-mono/configs"
	"go-mono/router"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// func handler(c *dandelion.Context) error {

// 	req := c.Request()
// 	fmt.Println(c, req)
// 	return c.JSON()
// }

func init() {
	configs.LoadEnv()
}

func main() {
	configs.ConnectDB("mysql")
	// echo version
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Static("/static", "static")
	router.Route(*e)

	const PORT = ":3001"
	e.Logger.Fatal(e.Start(PORT))

	// experimental feature
	// r := dandelion.NewRouter()
	// r.Route("GET", "/test", handler)

	// r.Route("GET", "/test2",
	// 	func(w http.ResponseWriter, r *http.Request) {
	// 		message := "test"
	// 		w.Write([]byte("Hello " + message))
	// 	})
	// http.ListenAndServe(":8000", dr)
}
