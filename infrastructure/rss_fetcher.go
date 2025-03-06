package infrastructure

import (
	"log"
	"time"

	"github.com/jt00721/daily-learning-feed/internal/entity"
	"github.com/mmcdole/gofeed"
)

type RSSFetcher struct {
	Parser *gofeed.Parser
}

func NewRSSFetcher() *RSSFetcher {
	return &RSSFetcher{
		Parser: gofeed.NewParser(),
	}
}

func (f *RSSFetcher) FetchResources(feedURL string) ([]entity.Resource, error) {
	feed, err := f.Parser.ParseURL(feedURL)
	if err != nil {
		log.Println("Failed to fetch RSS feed:", err)
		return nil, err
	}

	var resources []entity.Resource
	for _, item := range feed.Items {
		resource := entity.Resource{
			Title:     item.Title,
			URL:       item.Link,
			Category:  "RSS",
			Source:    feed.Title,
			DateAdded: time.Now(),
		}
		resources = append(resources, resource)
	}

	return resources, nil
}
