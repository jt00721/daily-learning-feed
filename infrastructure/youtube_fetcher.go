package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jt00721/daily-learning-feed/internal/domain"
)

type YouTubeFetcher struct {
	Client *resty.Client
	apiURL string // Add apiURL to allow test overrides
}

func NewYouTubeFetcher() *YouTubeFetcher {
	return &YouTubeFetcher{
		Client: resty.New(),
		apiURL: "https://www.googleapis.com/youtube/v3", // Default real API
	}
}

func (f *YouTubeFetcher) FetchVideos(query string) ([]domain.Resource, error) {
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		log.Println("Missing YOUTUBE_API_KEY")
		return nil, fmt.Errorf("missing API key")
	}

	url := fmt.Sprintf("%s/search", f.apiURL) // Use apiURL instead of hardcoded URL

	resp, err := f.Client.R().
		SetQueryParams(map[string]string{
			"part":       "snippet",
			"q":          query,
			"maxResults": "5",
			"type":       "video",
			"key":        apiKey,
		}).
		SetHeader("Accept", "application/json").
		Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch videos: %s", resp.Status())
	}

	var result struct {
		Items []struct {
			ID struct {
				VideoID string `json:"videoId"`
			} `json:"id"`
			Snippet struct {
				Title string `json:"title"`
			} `json:"snippet"`
		} `json:"items"`
	}

	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}

	if len(result.Items) == 0 {
		return nil, fmt.Errorf("no videos found for query: %s", query)
	}

	var resources []domain.Resource
	for _, item := range result.Items {
		resources = append(resources, domain.Resource{
			Title:     item.Snippet.Title,
			URL:       fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.ID.VideoID),
			Category:  "Video",
			Source:    "YouTube",
			DateAdded: time.Now(),
		})
	}

	return resources, nil
}
