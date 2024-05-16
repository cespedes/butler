package main

import (
	"errors"
	"fmt"
)

type Command func(args []string) error

var commands = make(map[string]Command)

func Register(name string, cmd Command) error {
	if commands[name] != nil {
		return errors.New(fmt.Sprintf("command %q already registered", name))
	}
	commands[name] = cmd
	return nil
}
