package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("no .env file found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("PORT is required")
	}

	return &Config{
		Port: ":" + port,
	}
}
