package main

import (
	"go-mono/configs"
	"go-mono/generator/pkg"
)

func init() {
	configs.LoadEnv()
}

func main() {
	pkg.GormGen()
}
