package commands

import (
	"fmt"

	"github.com/goodmustache/pt/actions"
	"github.com/goodmustache/pt/commands/internal/flags"
)

// UserCommand list information about a user or the current user
type UserCommand struct {
	flags.CommonFlags
}

// Execute is the execution of the UserCommand
func (cmd UserCommand) Execute([]string) error {
	user, err := actions.GetUser(cmd.Alias, cmd.Username)
	if err != nil {
		return err
	}

	fmt.Println("Name:", user.Name)
	fmt.Println("Username:", user.Username)

	return nil
}
