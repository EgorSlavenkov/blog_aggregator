package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("agg err: %v", err)
	}

	fmt.Printf("Channel: %s\n", feed.Channel.Title)
	fmt.Printf("Description: %s\n", feed.Channel.Description)

	fmt.Println("\nArticles:")
	for _, item := range feed.Channel.Item {
		fmt.Printf("\nTitle: %s\n", item.Title)
		fmt.Printf("Link: %s\n", item.Link)
		fmt.Printf("Published: %s\n", item.PubDate)
		fmt.Printf("Description: %s\n", item.Description)
	}

	return nil
}
