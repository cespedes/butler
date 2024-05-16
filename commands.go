package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func init() {
	Register("echo", cmdEcho)
	Register("cat", cmdCat)
}

func cmdEcho(args []string) error {
	fmt.Println(strings.Join(args[1:], " "))
	return nil
}

func cmdCat(args []string) error {
	var err error
	for _, arg := range args[1:] {
		f, err := os.Open(arg)
		if err != nil {
			return err
		}
		_, err = io.Copy(os.Stdout, f)
		if err != nil {
			f.Close()
			return err
		}
		err = f.Close()
		if err != nil {
			return err
		}
	}
	return err
}
