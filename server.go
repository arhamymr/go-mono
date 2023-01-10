package main

import (
	"go-mono/router"
)

const PORT = ":3000"

func main() {
	e := router.InitRoute()
	e.Logger.Fatal(e.Start(PORT))
}
