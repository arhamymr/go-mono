package router

import (
	"go-mono/handlers"

	"github.com/labstack/echo"
)

func InitRoute() echo.Echo {
	e := *echo.New()

	e.Static("/static", "static")

	// auth
	auth := e.Group("/auth")
	auth.POST("/register", handlers.Register)
	auth.POST("/login", handlers.Login)

	// post
	post := e.Group("/posts")
	post.GET("", handlers.GetMany)
	post.GET("/:id", handlers.GetOne)
	post.POST("/:id", handlers.Update)
	post.DELETE("/:id", handlers.Delete)

	return e
}
