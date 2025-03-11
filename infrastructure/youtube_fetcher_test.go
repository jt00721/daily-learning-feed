package infrastructure

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var youtubeMockResponse = `{
	"items": [
		{
			"id": { "videoId": "12345" },
			"snippet": { "title": "Test Video" }
		},
		{
			"id": { "videoId": "67890" },
			"snippet": { "title": "Another Video" }
		}
	]
}`

func TestFetchVideos(t *testing.T) {
	// Mock YouTube API response using httptest
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(youtubeMockResponse)) // Write mock response
	}))
	defer mockServer.Close()

	// Set mock API key
	os.Setenv("YOUTUBE_API_KEY", "testkey")

	// Override the fetcher to use the test server URL
	fetcher := NewYouTubeFetcher()
	fetcher.apiURL = mockServer.URL // Inject test server URL

	// Fetch mock data
	resources, err := fetcher.FetchVideos("golang")

	assert.NoError(t, err, "Expected no error when fetching mock data")
	assert.NotEmpty(t, resources, "Expected videos but got empty response")
	assert.Len(t, resources, 2, "Expected 2 videos but got a different count")
	assert.Equal(t, "Test Video", resources[0].Title)
	assert.Equal(t, "https://www.youtube.com/watch?v=12345", resources[0].URL)
}

func TestFetchVideos_MissingAPIKey(t *testing.T) {
	os.Setenv("YOUTUBE_API_KEY", "")

	fetcher := NewYouTubeFetcher()

	_, err := fetcher.FetchVideos("golang")

	assert.Error(t, err)
}

func TestFetchVideos_APIFailure(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))

	defer mockServer.Close()

	os.Setenv("YOUTUBE_API_KEY", "testkey")

	fetcher := NewYouTubeFetcher()
	fetcher.apiURL = mockServer.URL

	_, err := fetcher.FetchVideos("golang")
	assert.Error(t, err)
}
