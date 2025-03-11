package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jt00721/daily-learning-feed/internal/domain"
	"github.com/jt00721/daily-learning-feed/internal/repository"
)

type ResourceHandler struct {
	Repo *repository.ResourceRepository
}

func (h *ResourceHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("web/layout.html", "web/index.html")
	resources, _ := h.Repo.GetAll()

	data := struct {
		Title     string
		Resources []domain.Resource
	}{
		Title:     "Daily Learning Feed",
		Resources: resources,
	}

	tmpl.Execute(w, data)
}

func (h *ResourceHandler) AddResourcePage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("web/layout.html", "web/add_resource.html")
	tmpl.Execute(w, struct{ Title string }{Title: "Add Resource"})
}

func (h *ResourceHandler) CreateResourceForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	resource := domain.Resource{ // Ensure this uses the `entity.Resource` struct
		Title:    r.FormValue("title"),
		URL:      r.FormValue("url"),
		Category: r.FormValue("category"),
		Source:   r.FormValue("source"),
	}

	h.Repo.Create(&resource)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *ResourceHandler) CreateResourceJSON(c *gin.Context) {
	var resource domain.Resource
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
	var resource domain.Resource
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
