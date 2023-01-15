package main

import (
	"go-mono/configs"
	generator "go-mono/generator/pkg"
)

func init() {
	configs.LoadEnv()
}

func main() {
	generator.GormGen()
}
