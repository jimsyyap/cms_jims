package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jimsyyap/cms-project/internal/api"
	"github.com/jimsyyap/cms-project/internal/config"
	"github.com/jimsyyap/cms-project/internal/database"
	"github.com/jimsyyap/cms-project/internal/middleware"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database connection
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Set up Gin router
	router := gin.Default()

	// Apply middleware
	router.Use(middleware.CORS())
	router.Use(middleware.Security())

	// Set up API routes
	api.SetupRoutes(router, db)

	// Start the server
	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

