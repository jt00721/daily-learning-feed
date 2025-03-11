package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jt00721/daily-learning-feed/infrastructure"
	"github.com/jt00721/daily-learning-feed/internal/handler"
	"github.com/jt00721/daily-learning-feed/internal/repository"
)

// SetupRoutes defines all application routes
func SetupRoutes(router *gin.Engine, repo *repository.ResourceRepository) {
	rssFetcher := infrastructure.NewRSSFetcher()
	youtubeFetcher := infrastructure.NewYouTubeFetcher()
	devToFetcher := infrastructure.NewDevToFetcher()

	rssHandler := &handler.RSSHandler{Fetcher: rssFetcher, Repo: repo}
	resourceHandler := &handler.ResourceHandler{Repo: repo}
	apiHandler := &handler.APIHandler{
		YouTubeFetcher: youtubeFetcher,
		DevToFetcher:   devToFetcher,
		Repo:           repo,
	}

	router.GET("/", gin.WrapF(resourceHandler.HomePage))
	router.GET("/add", gin.WrapF(resourceHandler.AddResourcePage))
	router.POST("/resources", gin.WrapF(resourceHandler.CreateResourceForm))
	router.POST("/api/resources", resourceHandler.CreateResourceJSON)
	router.GET("/resources", resourceHandler.GetResources)
	router.GET("/resources/:id", resourceHandler.GetResourceByID)
	router.PUT("/resources/:id", resourceHandler.UpdateResource)
	router.DELETE("/resources/:id", resourceHandler.DeleteResource)

	router.GET("/fetch-rss", rssHandler.FetchAndStoreResources)
	router.GET("/fetch-youtube", apiHandler.FetchYouTubeVideos)
	router.GET("/fetch-devto", apiHandler.FetchDevToArticles)
}
