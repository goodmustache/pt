package commands

import (
	"fmt"

	"github.com/goodmustache/pt/actions"
	"github.com/goodmustache/pt/commands/internal/flags"
	"github.com/vito/go-interact/interact"
)

// RemoveUserCommand removes user from pt
type RemoveUserCommand struct {
	flags.CommonFlags
	Force bool `long:"with-malice" description:"Remove target user without confirmation"`
}

// Execute is the execution of the RemoveUserCommand
func (cmd RemoveUserCommand) Execute([]string) error {
	selectedUser, err := actions.GetUser(cmd.Alias, cmd.Username)
	if err != nil {
		return err
	}

	if !cmd.Force {
		message := fmt.Sprintf("Remove %s (%s):", selectedUser.Name, selectedUser.Username)
		var removeUser bool
		err := interact.NewInteraction(message).Resolve(interact.Required(&removeUser))
		if err != nil {
			return err
		}
		if !removeUser {
			return ErrUserTerminated
		}
	}

	err = actions.RemoveUser(selectedUser)
	if err != nil {
		return err
	}

	fmt.Printf("User %s (%s) has been removed.\n", selectedUser.Name, selectedUser.Username)

	return nil
}
