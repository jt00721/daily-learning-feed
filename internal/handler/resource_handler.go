package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jt00721/daily-learning-feed/internal/entity"
	"github.com/jt00721/daily-learning-feed/internal/repository"
)

type ResourceHandler struct {
	Repo *repository.ResourceRepository
}

func (h *ResourceHandler) CreateResource(c *gin.Context) {
	var resource entity.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.Repo.Create(&resource)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save resource"})
		return
	}

	c.JSON(http.StatusCreated, resource)
}

func (h *ResourceHandler) GetResources(c *gin.Context) {
	resources, err := h.Repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve resources"})
		return
	}

	c.JSON(http.StatusOK, resources)
}

func (h *ResourceHandler) GetResourceByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	resource, err := h.Repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		return
	}

	c.JSON(http.StatusOK, resource)
}

func (h *ResourceHandler) UpdateResource(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var resource entity.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	resource.ID = uint(id)
	h.Repo.Update(&resource)
	c.JSON(http.StatusOK, resource)
}

func (h *ResourceHandler) DeleteResource(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.Repo.Delete(uint(id))
	c.JSON(http.StatusOK, gin.H{"message": "Resource deleted"})
}
