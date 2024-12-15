package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PGHost     string
	PGPort     string
	PGUser     string
	PGPassword string
	PGDatabase string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		PGHost:     getEnv("PGHOST"),
		PGPort:     getEnv("PGPORT"),
		PGUser:     getEnv("PGUSER"),
		PGPassword: getEnv("PGPASSWORD"),
		PGDatabase: getEnv("PGDATABASE"),
	}
}

func getEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}

	return val
}