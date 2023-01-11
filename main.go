package main

import (
	"go-mono/configs"
	"go-mono/router"
)

const PORT = ":3000"

func init() {
	configs.LoadEnv()
}

func main() {
	r := router.InitRoute()
	r.Static("/static", "static")
	router.Route(*r)
	r.Logger.Fatal(r.Start(PORT))
}
