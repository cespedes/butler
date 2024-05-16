package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/chzyer/readline"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func run() error {
	rl, err := readline.New("> ")
	if err != nil {
		return err
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF
			break
		}
		fields := strings.Fields(line)

		if len(fields) == 0 {
			continue
		}
		cmd := commands[fields[0]]
		if cmd == nil {
			fmt.Fprintf(os.Stderr, "%s: command not found\n", fields[0])
			continue
		}
		err = cmd(fields)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", fields[0], err.Error())
			continue
		}
	}
	return nil
}
