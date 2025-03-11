package main

import (
	"github.com/jt00721/daily-learning-feed/config"
	"github.com/jt00721/daily-learning-feed/internal/routes"
)

func main() {
	// Initialize application
	application := config.NewApp()

	// Setup routes
	routes.SetupRoutes(application.Router, application.ResourceRepo)

	// Run application
	application.Run()
}
