package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application's configuration settings.
type Config struct {
	MongoDBURI string
	JWTSecret  string
	Port       string
}

// LoadConfig loads configuration from environment variables or .env file.
func LoadConfig() *Config {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables.")
	}

	cfg := &Config{
		MongoDBURI: os.Getenv("MONGODB_URI"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		Port:       os.Getenv("PORT"),
	}

	if cfg.MongoDBURI == "" {
		log.Fatal("MONGODB_URI not set in environment variables or .env file")
	}
	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET not set in environment variables or .env file")
	}
	if cfg.Port == "" {
		cfg.Port = "8080" // Default port
	}

	return cfg
}
