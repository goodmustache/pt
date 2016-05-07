package commands

import (
	"fmt"

	"github.com/goodmustache/pt/actions"
)

type UserCommand struct {
	Alias    string `short:"a" long:"alias" description:"Target user's alias"`
	Username string `short:"u" long:"username" description:"Target user's username"`
}

func (cmd *UserCommand) Execute([]string) error {
	user, err := actions.GetUser(cmd.Alias, cmd.Username)
	if err != nil {
		return err
	}

	fmt.Println("Name:", user.Name)
	fmt.Println("Username:", user.Username)

	return nil
}
