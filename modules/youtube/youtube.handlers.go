package mod_youtubes

import (
	"go-mono/apis"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SearchYoutube(c echo.Context) error {
	result := &apis.FinalResult{}
	keyword := "this keyword"
	err := GetYoutubeVideo(keyword, result)

	if err != nil {
		panic("failed to load data")
	}

	return c.String(http.StatusOK, "hello login")
}
