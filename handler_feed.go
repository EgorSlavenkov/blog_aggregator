package main

import (
	"context"
	"fmt"
	"time"

	"github.com/EgorSlavenkov/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handleFeeds(s *state, cmd command) error {
	feedsList, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("cant get feeds list: %v", err)
	}
	fmt.Println(feedsList)
	return nil
}

func handleAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting current user: %w", err)
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	fmt.Printf("Feed %s created successfully!\n", feed.Name)
	return nil
}
