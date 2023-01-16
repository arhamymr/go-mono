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

func CustomException(err error) *echo.HTTPError {

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

	notFound := strings.Contains(err.Error(), "not found")

	if notFound {
		return echo.NewHTTPError(
			http.StatusNotFound,
			ErrorMessage{
				Status:  http.StatusBadRequest,
				Message: "Data not found",
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
