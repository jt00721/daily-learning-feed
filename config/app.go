package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jt00721/daily-learning-feed/infrastructure"
	"github.com/jt00721/daily-learning-feed/internal/repository"
)

// App holds application dependencies
type App struct {
	Router         *gin.Engine
	ResourceRepo   *repository.ResourceRepository
	RSSFetcher     *infrastructure.RSSFetcher
	YouTubeFetcher *infrastructure.YouTubeFetcher
	DevToFetcher   *infrastructure.DevToFetcher
}

// NewApp initializes the app
func NewApp() *App {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	infrastructure.InitDB()

	// Initialize repositories & fetchers
	resourceRepo := &repository.ResourceRepository{DB: infrastructure.DB}
	rssFetcher := infrastructure.NewRSSFetcher()
	youtubeFetcher := infrastructure.NewYouTubeFetcher()
	devToFetcher := infrastructure.NewDevToFetcher()

	// Create Gin router
	router := gin.Default()
	router.Static("/static", "./static")

	return &App{
		Router:         router,
		ResourceRepo:   resourceRepo,
		RSSFetcher:     rssFetcher,
		YouTubeFetcher: youtubeFetcher,
		DevToFetcher:   devToFetcher,
	}
}

// Run starts the server
func (app *App) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port", port)
	app.Router.Run(":" + port)
}
