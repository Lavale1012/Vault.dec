package main

import (
	"fmt"
	"os"
	"vault-dev/config"
	"vault-dev/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading environment variables:", err)
		return
	}

	PORT := os.Getenv("PORT")

	// Create a new Gin router
	r := gin.Default()
	r.Use(cors.Default())
	// Set up the routes
	routes.Snippet_Routes(r)
	config.ConnectDB()

	if err := r.Run("localhost:" + PORT); err != nil {
		fmt.Println("Failed to start server:", err)
	}

}

// This is the main entry point for the Go application.
