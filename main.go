package main

import (
	"go-mono/configs"
	"go-mono/router"

	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

const PORT = ":3000"

func init() {
	configs.LoadEnv()
}

func main() {
	r := router.InitRoute()
	r.GET("/docs/*", echoSwagger.WrapHandler)
	r.Logger.Fatal(r.Start(PORT))
}
