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
			usageHelp()
		}
		os.Exit(1)
	}
}

func usageHelp() {
	fmt.Fprintf(os.Stderr, "Usage: %s {list|connect|help}\n", os.Args[0])
}

func run() error {
	flags, err := config()
	if err != nil {
		return fmt.Errorf("config error: %v\n", err)
	}
	switch cmd := flags.Arg(0); cmd {
	case "list":
		return cmdList()
	case "help":
		flags.Usage()
		return nil
	default:
		return unknownCommandError{cmd: cmd}
	}
}
