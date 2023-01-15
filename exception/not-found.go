package exception

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type ErrorMessage struct {
	Status  uint
	Message string
	Detail  string
}

func NotFound(err error) *echo.HTTPError {

	duplicateEntry := strings.Contains(err.Error(), "Duplicate entry")
	if duplicateEntry {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			ErrorMessage{
				Status:  http.StatusBadRequest,
				Message: "Duplicate Entry, Data Already Exists",
				Detail:  err.Error(),
			})
	}
	return echo.NewHTTPError(
		http.StatusNotImplemented,
		ErrorMessage{
			Status:  http.StatusNotImplemented,
			Message: "Something Wrong",
			Detail:  err.Error(),
		})
}
