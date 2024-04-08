package main

import (
	"backend-service/router"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}
	r := router.SetupRouter()

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Run(":8080")
}
