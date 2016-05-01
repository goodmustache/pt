package main

import (
	"fmt"
	"os"

	"github.com/goodmustache/pt/commands"
	"github.com/jessevdk/go-flags"
)

func main() {
	parser := flags.NewParser(&commands.PT, flags.HelpFlag)
	parser.NamespaceDelimiter = "-"

	_, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
