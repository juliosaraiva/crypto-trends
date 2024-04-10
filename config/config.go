package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
	// load .env file
	_ = godotenv.Load(".env")
	return os.Getenv(key)
}