package router

import (
	"go-mono/handlers"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitRoute() *echo.Echo {
	return echo.New()
}

func Route(r echo.Echo) {
	// auth
	auth := r.Group("/auth", middleware.Secure(), middleware.Logger())
	auth.POST("/register", handlers.Register)
	auth.POST("/login", handlers.Login)

	// post
	post := r.Group("/posts", middleware.Secure(), middleware.Logger())
	post.GET("", handlers.GetMany)
	post.GET("/:id", handlers.GetOne)
	post.DELETE("/:id", handlers.Delete)

	// docs
	r.GET("/docs/*", echoSwagger.WrapHandler, middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().URL.Path, "docs") {
				return true
			}
			return false
		},
	}))
}
