package router

import (
	"go-mono/jwt"
	mod_auth "go-mono/modules/auth"
	mod_media "go-mono/modules/media"
	mod_posts "go-mono/modules/posts"
	"strings"

	echojwt "github.com/labstack/echo-jwt/v4"
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
	auth.POST("/register", mod_auth.Register)
	auth.POST("/login", mod_auth.Login)
	// todo

	// delete users
	// update users

	// post
	post := r.Group("/post", middleware.Secure())

	post.Use(echojwt.WithConfig(jwt.Configs()))
	post.POST("/create", mod_posts.Create)
	post.GET("", mod_posts.GetMany)
	post.GET("/:id", mod_posts.GetOne)
	post.DELETE("/:id", mod_posts.Delete)

	// media
	media := r.Group("/media")
	media.POST("/upload", mod_media.Upload)
	// todo
	media.DELETE("/delete", mod_media.Delete)

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
