package main

import (
	"cmp"
	"fmt"
	"io"
	"os"
	"strings"
)

func init() {
	Register("echo", cmdEcho)
	Register("cat", cmdCat)
	Register("cd", cmdCd)
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

func cmdCd(args []string) error {
	if len(args) == 1 {
		return os.Chdir(cmp.Or(os.Getenv("HOME"), "/"))
	}
	if len(args) > 2 {
		return fmt.Errorf("wrong number of arguments")
	}
	return os.Chdir(args[1])
}
