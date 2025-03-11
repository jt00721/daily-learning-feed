package infrastructure

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock Dev.to API response
var devtoMockResponse = `[
	{ "title": "Golang Tips", "url": "https://dev.to/golang-tips" },
	{ "title": "REST API with Go", "url": "https://dev.to/rest-api-go" }
]`

// Test FetchArticles()
func TestFetchArticles(t *testing.T) {
	// Mock Dev.to API response using httptest
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(devtoMockResponse)) // Write mock response
	}))
	defer mockServer.Close()

	// Override the fetcher to use the test server URL
	fetcher := NewDevToFetcher()
	fetcher.apiURL = mockServer.URL // Inject test server URL

	// Fetch mock data
	resources, err := fetcher.FetchArticles()

	assert.NoError(t, err, "Expected no error when fetching mock data")
	assert.NotEmpty(t, resources, "Expected articles but got empty response")
	assert.Len(t, resources, 2, "Expected 2 articles but got a different count")
	assert.Equal(t, "Golang Tips", resources[0].Title)
	assert.Equal(t, "https://dev.to/golang-tips", resources[0].URL)
}

func TestFetchArticles_ApiFailure(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
	}))
	defer mockServer.Close()

	fetcher := NewDevToFetcher()
	fetcher.apiURL = mockServer.URL

	_, err := fetcher.FetchArticles()
	assert.Error(t, err) // Should return an error
}
