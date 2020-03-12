package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/goodmustache/pt/actor"
	"github.com/goodmustache/pt/command"
	"github.com/goodmustache/pt/config"
	"github.com/goodmustache/pt/tracker"
	"github.com/goodmustache/pt/ui"
	flags "github.com/jessevdk/go-flags"
)

var uiWrapper *ui.UI

func main() {
	uiWrapper = ui.NewUI()

	parser := flags.NewParser(&command.PT, flags.HelpFlag)
	parser.CommandHandler = commandWrapper

	_, err := parser.Parse()
	switch {
	case err == nil:
	default:
		uiWrapper.PrintError(err)
		os.Exit(1)
	}
}

func commandWrapper(cmd flags.Commander, args []string) error {
	if len(args) > 0 {
		uiWrapper.PrintWarning("ignoring extra args: %s", strings.Join(args, " "))
	}

	config.SetDefaultConfig()
	cfg, err := config.ReadConfig()
	if err != nil {
		return fmt.Errorf("Error reading config file: %w", err)
	}

	switch t := cmd.(type) {
	case *command.UserAdd:
		client := tracker.NewClient(tracker.Config{APIToken: t.APIToken})
		t.Actor = actor.NewActor(client, cfg)
		t.Config = cfg
		t.UI = uiWrapper
	case *command.UserList:
		t.Config = cfg
		t.UI = uiWrapper
	}

	err = cmd.Execute(args)
	if err != nil {
		return err
	}

	err = config.WriteConfig()
	if err != nil {
		return fmt.Errorf("Error writing config file: %w", err)
	}

	return nil
}
