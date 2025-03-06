package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jt00721/daily-learning-feed/infrastructure"
	"github.com/jt00721/daily-learning-feed/internal/repository"
)

type RSSHandler struct {
	Fetcher *infrastructure.RSSFetcher
	Repo    *repository.ResourceRepository
}

func (h *RSSHandler) FetchAndStoreResources(c *gin.Context) {
	feedURL := c.Query("url")

	if feedURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing RSS feed URL"})
		return
	}

	resources, err := h.Fetcher.FetchResources(feedURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch RSS feed"})
		return
	}

	for _, resource := range resources {
		h.Repo.Create(&resource)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Resources fetched and saved", "count": len(resources)})
}
