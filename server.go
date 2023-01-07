package main

import (
	"go-mono/morestrings"
	"go-mono/word"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	const PORT = ":1234"
	router := gin.Default()
	router.Static("/static", "./static")
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Welcome to gin",
			"author":  word.Hello("arham"),
			"reverse": morestrings.ReverseRunes(word.Hello("arham")),
		})
	})
	router.Run(PORT)
}
