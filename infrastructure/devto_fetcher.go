package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jt00721/daily-learning-feed/internal/domain"
)

type DevToFetcher struct {
	Client *resty.Client
	apiURL string // Add apiURL to allow test overrides
}

func NewDevToFetcher() *DevToFetcher {
	return &DevToFetcher{
		Client: resty.New(),
		apiURL: "https://dev.to/api", // Default real API
	}
}

func (f *DevToFetcher) FetchArticles() ([]domain.Resource, error) {
	url := fmt.Sprintf("%s/articles", f.apiURL) // Use apiURL instead of hardcoded URL

	resp, err := f.Client.R().
		SetHeader("Accept", "application/json").
		Get(url)

	if err != nil {
		log.Println("Error fetching Dev.to articles:", err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch articles: %s", resp.Status())
	}

	var result []struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	}

	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}

	const maxArticles = 5
	if len(result) > maxArticles {
		result = result[:maxArticles]
	}

	var resources []domain.Resource
	for _, item := range result {
		resources = append(resources, domain.Resource{
			Title:     item.Title,
			URL:       item.URL,
			Category:  "Article",
			Source:    "Dev.to",
			DateAdded: time.Now(),
		})
	}

	return resources, nil
}
