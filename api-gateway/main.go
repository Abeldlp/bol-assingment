package main

import (
	"os"

	"github.com/Abeldlp/bol-assingment/api-gateway/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("PROD_ENV")

	if env == "" {
		godotenv.Load(".env")
	}

	r := gin.Default()

	r.Any("/v1/api/*proxyPath", handlers.ProxyTo(os.Getenv("API_SERVICE_URL")))

	r.Run(":5000")
}
