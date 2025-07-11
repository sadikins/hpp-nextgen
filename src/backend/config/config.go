package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadConfig loads environment variables from a .env file.
func LoadConfig() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	// Construct the absolute path to the .env file in the project root
	envPath := filepath.Join(cwd, "..", "..", ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file from %s: %v", envPath, err)
	}
}