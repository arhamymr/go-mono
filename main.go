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
	e := router.InitRoute()
	e.Logger.Fatal(e.Start(PORT))
}
