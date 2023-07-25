package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadENV(path string) {
	fmt.Println("path", path)
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appName := os.Getenv("APP_NAME")
	fmt.Println("appName", appName)
}
