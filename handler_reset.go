package main

import (
	"context"
	"fmt"
	"os"
)

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
