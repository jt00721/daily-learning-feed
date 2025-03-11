package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	domain "github.com/jt00721/daily-learning-feed/internal/domain"
	"github.com/jt00721/daily-learning-feed/internal/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestServer() (*gin.Engine, *repository.ResourceRepository) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&domain.Resource{})

	repo := &repository.ResourceRepository{DB: db}
	handler := &ResourceHandler{Repo: repo}

	r := gin.Default()
	r.POST("/resources", handler.CreateResourceJSON)
	r.GET("/resources", handler.GetResources)
	r.GET("/resources/:id", handler.GetResourceByID)
	r.PUT("/resources/:id", handler.UpdateResource)
	r.DELETE("/resources/:id", handler.DeleteResource)

	return r, repo
}

func TestCreateResourceAPI(t *testing.T) {
	router, _ := setupTestServer()

	newResource := domain.Resource{Title: "Golang Docs", URL: "https://golang.org"}
	jsonBody, _ := json.Marshal(newResource)

	req, _ := http.NewRequest("POST", "/resources", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetAllResourcesAPI(t *testing.T) {
	router, repo := setupTestServer()

	repo.Create(&domain.Resource{Title: "Go Basics", URL: "https://golang.org"})
	repo.Create(&domain.Resource{Title: "Gin Framework", URL: "https://gin-gonic.com"})

	req, _ := http.NewRequest("GET", "/resources", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateResourceAPI(t *testing.T) {
	router, repo := setupTestServer()

	resource := domain.Resource{Title: "To Delete", URL: "https://delete.com"}
	repo.Create(&resource)

	updatedResource := domain.Resource{Title: "Updated Title", URL: "https://updated.com"}
	jsonBody, _ := json.Marshal(updatedResource)

	path := fmt.Sprintf("/resources/%d", resource.ID)

	req, _ := http.NewRequest("PUT", path, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteResourceAPI(t *testing.T) {
	router, repo := setupTestServer()

	resource := domain.Resource{Title: "To Delete", URL: "https://delete.com"}
	repo.Create(&resource)

	path := fmt.Sprintf("/resources/%d", resource.ID)

	req, _ := http.NewRequest("DELETE", path, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
