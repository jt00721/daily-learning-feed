package repository

import (
	"testing"

	"github.com/jt00721/daily-learning-feed/internal/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&domain.Resource{})
	return db
}

func TestCreateResource(t *testing.T) {
	db := setupTestDB()
	repo := &ResourceRepository{DB: db}

	resource := domain.Resource{Title: "Test Title", URL: "https://test.com"}
	err := repo.Create(&resource)

	assert.NoError(t, err)
	assert.NotZero(t, resource.ID)
}

func TestGetAllResources(t *testing.T) {
	db := setupTestDB()
	repo := &ResourceRepository{DB: db}

	repo.Create(&domain.Resource{Title: "Golang Guide", URL: "https://golang.org"})
	repo.Create(&domain.Resource{Title: "Gin Framework", URL: "https://gin-gonic.com"})

	resources, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, resources, 2)
}

func TestGetByID(t *testing.T) {
	db := setupTestDB()
	repo := &ResourceRepository{DB: db}

	resource := domain.Resource{Title: "Test", URL: "https://example.com"}
	repo.Create(&resource)

	fetchedResource, err := repo.GetByID(resource.ID)
	assert.NoError(t, err)
	assert.Equal(t, resource.Title, fetchedResource.Title)
}

func TestUpdateResource(t *testing.T) {
	db := setupTestDB()
	repo := &ResourceRepository{DB: db}

	resource := domain.Resource{Title: "Update Me", URL: "https://update.com"}
	repo.Create(&resource)

	updatedResource := domain.Resource{ID: resource.ID, Title: "Updated Title", URL: "https://updated-resource.com"}
	err := repo.Update(&updatedResource)
	assert.NoError(t, err)

	fetchedResource, _ := repo.GetByID(resource.ID)
	assert.Equal(t, updatedResource.Title, fetchedResource.Title)
}

func TestDeleteResource(t *testing.T) {
	db := setupTestDB()
	repo := &ResourceRepository{DB: db}

	resource := domain.Resource{Title: "Delete Me", URL: "https://delete.com"}
	repo.Create(&resource)

	err := repo.Delete(resource.ID)
	assert.NoError(t, err)

	_, err = repo.GetByID(resource.ID)
	assert.Error(t, err)
}
