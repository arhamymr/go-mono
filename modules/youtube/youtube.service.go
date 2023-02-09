package mod_youtubes

import (
	"go-mono/apis"
)

func PostService() string {
	return "service posts"
}

func GetYoutubeVideo(keyword string, result *apis.FinalResult) error {
	err := apis.SearchVideos(keyword, result)
	return err
}
