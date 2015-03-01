package main

import (
	"fmt"
	"os"

	flag "github.com/ogier/pflag"
)

func config() (*flag.FlagSet, error) {
	flags := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fmt.Print("config()\n")
	return flags, nil
}
