package main

import (
	"fmt"
	"os"
	"os/exec"
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
		c := commands[fields[0]]
		if c != nil {
			err = c(fields)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %s\n", fields[0], err.Error())
			}
			continue
		}
		cmd := exec.Command(fields[0], fields[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", fields[0], err.Error())
		}
		continue
	}
	return nil
}
