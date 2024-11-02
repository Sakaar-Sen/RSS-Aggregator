package main

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Sakaar-Sen/rssagg/internal/database"
	"github.com/google/uuid"
)

func startScraping(db *database.Queries, concurrency int, timebetweenrequest time.Duration) {
	log.Printf("Starting scraping with %v workers with time between requests of %v", concurrency, timebetweenrequest)
	ticker := time.NewTicker(timebetweenrequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Printf("Error fetching feeds: %v", err)
			continue
		}

		if len(feeds) == 0 {
			log.Printf("No feeds to fetch")
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}

		wg.Wait()

	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed fetched: %v", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed: %v", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		t, err := time.Parse(time.RFC1123Z, item.PubDate)

		if err != nil {
			log.Printf("Error parsing time: %v", err)
			continue
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Description: item.Description,
			PublishedAt: t,
			Url:         item.Link,
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "unique constraint") {
				continue
			}
			log.Printf("Error creating post: %v", err)
		}
	}

	log.Printf("Fetched feed: %v", feed.Url)
}
