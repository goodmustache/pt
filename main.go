package main

import (
	"fmt"
	"os"

	"github.com/goodmustache/pt/command"
	flags "github.com/jessevdk/go-flags"
)

func main() {
	parser := flags.NewParser(&command.PT, flags.HelpFlag)

	_, err := parser.Parse()
	switch {
	case err == nil:
	default:
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
