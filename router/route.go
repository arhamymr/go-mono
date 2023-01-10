package router

import (
	"go-mono/handlers"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute() echo.Echo {
	r := *echo.New()

	// middleware

	// skip swagger docs
	r.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().URL.Path, "docs") {
				return true
			}
			return false
		},
	}))

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
