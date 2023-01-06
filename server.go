package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    const PORT = ":1234"
    router := gin.Default()
    
    // serving static 
    router.Static("/static", "./static")
    
    router.GET("/",func(context *gin.Context) {
        context.JSON(http.StatusOK, gin.H{
            "message": "Welcome to gin",
        })
    }) 
   router.Run(PORT)
}