package router

import (
	pkg_auth "go-mono/pkg/auth"
	pkg_posts "go-mono/pkg/posts"
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
	auth := r.Group("/auth", middleware.Secure())
	auth.POST("/register", pkg_auth.Register)
	auth.POST("/login", pkg_auth.Login)

	// post
	post := r.Group("/posts", middleware.Secure())
	post.GET("", pkg_posts.GetMany)
	post.GET("/:id", pkg_posts.GetOne)
	post.DELETE("/:id", pkg_posts.Delete)

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
