package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jt00721/daily-learning-feed/infrastructure"
	"github.com/jt00721/daily-learning-feed/internal/handler"
	"github.com/jt00721/daily-learning-feed/internal/repository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	infrastructure.InitDB()

	r := gin.Default()

	repo := &repository.ResourceRepository{DB: infrastructure.DB}
	rssFetcher := infrastructure.NewRSSFetcher()
	rssHandler := &handler.RSSHandler{Fetcher: rssFetcher, Repo: repo}
	resourceHandler := &handler.ResourceHandler{Repo: repo}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the Daily Learning Feed!"})
	})

	r.POST("/resources", resourceHandler.CreateResource)
	r.GET("/resources", resourceHandler.GetResources)
	r.GET("/resources/:id", resourceHandler.GetResourceByID)
	r.PUT("/resources/:id", resourceHandler.UpdateResource)
	r.DELETE("/resources/:id", resourceHandler.DeleteResource)

	r.GET("/fetch-rss", rssHandler.FetchAndStoreResources)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port", port)
	r.Run(":" + port)
}
