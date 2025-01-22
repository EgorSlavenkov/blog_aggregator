package main

import (
	"fmt"
	"log"

	"github.com/EgorSlavenkov/blog_aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	cfg.SetUser("egor")
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("DB URL: %s, Current User: %s\n", cfg.DbURL, cfg.CurrentUserName)
}
