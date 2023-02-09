package apis

import (
	"flag"
	"fmt"
	"go-mono/configs"
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var maxResults = flag.Int64("max-results", 25, "Max YouTube results")

type FinalResult struct {
	videos *[]SearchResult
}

type SearchResult struct {
	Title       string `json:"title"`
	VideoId     string `json:"videoId"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
}

func SearchVideos(keyword string, result *FinalResult) error {
	flag.Parse()

	client := &http.Client{
		Transport: &transport.APIKey{Key: configs.GET("GCP_API_KEY")},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	snippet := []string{"id", "snippet"}

	call := service.Search.List(snippet).
		Q(keyword).
		MaxResults(*maxResults)
	response, err := call.Do()

	if err != nil {
		panic("error" + err.Error())
	}

	if err != nil {
		panic("failed marshal" + err.Error())
	}

	var searchResult []SearchResult

	// for _, value := range response.Items {
	// 	loopdata := SearchResult{
	// 		value.Snippet.Title,
	// 		value.Id.VideoId,
	// 		value.Snippet.Description,
	// 		value.Snippet.Thumbnails.Default.Url,
	// 	}
	// 	searchResult = append(searchResult, loopdata)
	// }

	result.videos = &searchResult
	fmt.Println(response, result, "result")
	return err
}
