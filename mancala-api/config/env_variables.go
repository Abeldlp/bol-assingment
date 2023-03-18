package config

import (
	"os"

	"github.com/joho/godotenv"
)

func InitializeEnvironmentVariables() {
	env := os.Getenv("PROD_ENV")
	if env == "" {
		godotenv.Load(".env")
	}
}
