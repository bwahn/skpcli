package main

import (
	"fmt"
	"os"

	flag "github.com/ogier/pflag"
)

func config() (*flag.FlagSet, error) {
	/*
		dir, err := configDir(".skpcli")
		if err != nil {
			return nil, fmt.Errorf("Failed to get .skpcli dir: %s", err)
		}*/
	flags := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flags.Usage = func() { usageHelp(flags) }

	if err := flags.Parse([]string{}); err != nil {
		return nil, err
	}

	return flags, nil
}

func usageHelp(flags *flag.FlagSet) {
	fmt.Fprintf(os.Stderr, `Usage: %s [<options>] <command> [<args>]

skpcli utility.

Commands:
	list      Display servers
	connect   Connect server

Options:
`, os.Args[0])
	flags.PrintDefaults()
}

func usageShort() {
	fmt.Fprintf(os.Stderr, "Usage: %s {list|connect|help} [<args>]\n", os.Args[0])
}
