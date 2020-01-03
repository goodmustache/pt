package main

import (
	"fmt"
	"os"

	"github.com/goodmustache/pt/command"
	"github.com/goodmustache/pt/config"
	flags "github.com/jessevdk/go-flags"
)

func main() {
	parser := flags.NewParser(&command.PT, flags.HelpFlag)
	parser.CommandHandler = commandWrapper

	_, err := parser.Parse()
	switch {
	case err == nil:
	default:
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func commandWrapper(cmd flags.Commander, args []string) error {
	config.SetDefaultConfig()
	_, err := config.ReadConfig()
	if err != nil {
		return err
	}

	err = cmd.Execute(args)
	if err != nil {
		return err
	}

	return config.WriteConfig()
}
