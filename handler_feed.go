package main

import (
	"context"
	"fmt"
	"time"

	"github.com/EgorSlavenkov/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

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
		UserID: uuid.NullUUID{
			UUID:  user.ID,
			Valid: true,
		},
	}

	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	fmt.Printf("Feed %s created successfully!\n", feed.Name)
	return nil
}
