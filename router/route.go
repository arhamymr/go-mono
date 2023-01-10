package router

import (
	"go-mono/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRoute() echo.Echo {
	r := *echo.New()

	// middleware
	r.Use(middleware.Gzip())
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())
	r.Use(middleware.Secure())

	r.Static("/static", "static")

	// auth
	auth := r.Group("/auth")
	auth.POST("/register", handlers.Register)
	auth.POST("/login", handlers.Login)

	// post
	post := r.Group("/posts")
	post.GET("", handlers.GetMany)
	post.GET("/:id", handlers.GetOne)
	post.POST("/:id", handlers.Update)
	post.DELETE("/:id", handlers.Delete)

	return r
}
