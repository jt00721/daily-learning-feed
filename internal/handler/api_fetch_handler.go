package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jt00721/daily-learning-feed/infrastructure"
	"github.com/jt00721/daily-learning-feed/internal/repository"
)

type APIHandler struct {
	YouTubeFetcher *infrastructure.YouTubeFetcher
	DevToFetcher   *infrastructure.DevToFetcher
	Repo           *repository.ResourceRepository
}

func (h *APIHandler) FetchYouTubeVideos(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing query parameter"})
		return
	}

	videos, err := h.YouTubeFetcher.FetchVideos(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch YouTube videos " + fmt.Sprintf("%v", err)})
		return
	}

	for _, video := range videos {
		h.Repo.Create(&video)
	}

	c.JSON(http.StatusOK, gin.H{"message": "YouTube videos fetched and saved", "count": len(videos)})
}

func (h *APIHandler) FetchDevToArticles(c *gin.Context) {
	articles, err := h.DevToFetcher.FetchArticles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Dev.to articles"})
		return
	}

	for _, article := range articles {
		h.Repo.Create(&article)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dev.to articles fetched and saved", "count": len(articles)})
}
