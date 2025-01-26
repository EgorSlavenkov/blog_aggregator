package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/EgorSlavenkov/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("user '%s' does not exist\n", name)
			os.Exit(1)
		}
		return fmt.Errorf("error checking user: %w", err)
	}
	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("no name provided for registration")
	}
	name := cmd.Args[0]
	currName, err := s.db.GetUser(context.Background(), name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("error querying user: %w", err)
	}
	if err == nil && currName.Name == name {
		os.Exit(1)
	}
	createParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}
	_, err = s.db.CreateUser(context.Background(), createParams)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	s.cfg.SetUser(name)
	fmt.Printf("successful user %v registration\n", name)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		fmt.Println("fail to delete all users")
		os.Exit(1)
	}
	fmt.Println("successfully deleted all users")
	os.Exit(0)
	return nil
}
