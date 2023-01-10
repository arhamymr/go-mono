package configs

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
}

func GET(key string) string {
	return os.Getenv(key)
}
