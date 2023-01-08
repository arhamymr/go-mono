package main

import (
	"fmt"
	"go-mono/configs"
	"go-mono/morestrings"
	"go-mono/word"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Register struct {
	email    string
	password string
	name     string
}

var db *gorm.DB

const PORT = ":3000"

func init() {
	db = configs.ConnectDB()
}

func main() {
	router := gin.Default()
	router.Static("/static", "./static")

	router.POST("/auth/register", func(c *gin.Context) {
		fmt.Println(c, "context")

		var reg Register

		if err := c.ShouldBind(&reg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "samothing wrong",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "created users",
			"payload": reg.email,
			"data":    reg.name,
			"all":     reg,
		})
	})

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Welcome to gin",
			"author":  word.Hello("arham"),
			"reverse": morestrings.ReverseRunes(word.Hello("arham")),
			// "result":  result,
		})
	})

	router.Run(PORT)
}
