package main

import (
	"fmt"
	"os"
)

type unknownCommandError struct {
	cmd string
}

func (e unknownCommandError) Error() string {
	return fmt.Sprintf("Unkown command: %s", e.cmd)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error in run: %v\n", err)
		if _, ok := err.(unknownCommandError); ok {
			usageShort()
		}
		os.Exit(1)
	}
}

func run() error {
	flags, err := config()
	if err != nil {
		return fmt.Errorf("config error: %v\n", err)
	}
	cmd := flags.Arg(0)
	fmt.Fprintf(os.Stderr, "cmd : %v\n", cmd)
	switch cmd {
	case "list":
		return cmdList()
	case "help":
		fmt.Fprintf(os.Stderr, "aaa\n")
		flags.Usage()
		return nil
	case "":
		fmt.Fprintf(os.Stderr, "bbb\n")
		usageShort()
		return nil
	default:
		return unknownCommandError{cmd: cmd}
	}
}
