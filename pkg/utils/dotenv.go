package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func InitDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
