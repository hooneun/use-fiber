package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .evn file")
	}
	return os.Getenv(key)
}
