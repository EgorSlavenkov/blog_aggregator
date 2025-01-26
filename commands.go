package main

import (
	"errors"
)

type command struct {
	name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	currFunc, ok := c.registeredCommands[cmd.name]
	if !ok {
		return errors.New("no command exist")
	}
	return currFunc(s, cmd)
}
