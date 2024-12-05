package system

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitLoadEnv() {
	if os.Getenv("GO_ENV") == "develop" || os.Getenv("GO_ENV") == "" {
		loadDotEnv()
	}
}

// Get variables from .env files.
func loadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}
