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
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
