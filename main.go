package main

import (
	"fmt"
	"os"

	"github.com/goodmustache/pt/commands"
	flags "github.com/jessevdk/go-flags"
)

func main() {
	parser := flags.NewParser(&commands.PT, flags.HelpFlag)

	_, err := parser.Parse()
	switch {
	case err == nil:
	case err == commands.ErrUserTerminated:
		os.Exit(1)
	default:
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
