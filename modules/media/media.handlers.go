package mod_media

import (
	"go-mono/exception"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Upload(c echo.Context) error {
	// if you what to extract token
	// token := jwt.ExtractToken(c)
	file, err := c.FormFile("file")

	if err != nil {
		return err
	}

	err = UploadMedia(file)

	if err != nil {
		return exception.CustomException(err)
	}

	return c.JSON(200, Result{
		Status:  http.StatusOK,
		Message: "File Uploaded " + file.Filename,
	})
}

func Delete(c echo.Context) error {
	return c.String(http.StatusOK, "hello login")
}
