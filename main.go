package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/EgorSlavenkov/blog_aggregator/internal/config"
	"github.com/EgorSlavenkov/blog_aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	dbURL := "postgres://postgres:123@localhost:5432/gator"
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handleAddFeed)
	cmds.register("feeds", handleFeeds)

	if len(os.Args) < 2 {
		log.Fatal("Usage: CLI <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
