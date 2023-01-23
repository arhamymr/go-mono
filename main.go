package main

import (
	"go-mono/configs"
	drout "go-mono/experimental"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

const PORT = ":3000"

func init() {
	configs.LoadEnv()
}

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

func main() {

	// disable echo for working on experiment feature
	// e := echo.New()
	// e.Validator = &CustomValidator{validator: validator.New()}
	// e.Static("/static", "static")
	// router.Route(*e)
	// e.Logger.Fatal(e.Start(PORT))

	// experimental feature
	dr := drout.NewRouter()
	dr.Route("GET", "/test",
		func(w http.ResponseWriter, r *http.Request) {
			message := "test"
			w.Write([]byte("Hello " + message))
		})

	dr.Route("GET", "/test2",
		func(w http.ResponseWriter, r *http.Request) {
			message := "test"
			w.Write([]byte("Hello " + message))
		})
	http.ListenAndServe(":8000", dr)
}
