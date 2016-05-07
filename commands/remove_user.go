package commands

import (
	"fmt"

	"github.com/goodmustache/pt/actions"
	"github.com/vito/go-interact/interact"
)

type RemoveUserCommand struct {
	Alias    string `short:"a" long:"alias" description:"Target user's alias"`
	Username string `short:"u" long:"username" description:"Target user's username"`
	Force    bool   `long:"with-malice" description:"Remove target user without confirmation"`
}

func (cmd *RemoveUserCommand) Execute([]string) error {
	selectedUser, err := actions.GetUser(cmd.Alias, cmd.Username)
	if err != nil {
		return err
	}

	if !cmd.Force {
		message := fmt.Sprintf("Remove %s (%s):", selectedUser.Name, selectedUser.Username)
		var input bool
		err := interact.NewInteraction(message).Resolve(interact.Required(&input))
		if err != nil {
			return err
		}
	}

	actions.RemoveUser(selectedUser)

	fmt.Printf("User %s (%s) has been removed.\n", selectedUser.Name, selectedUser.Username)

	return nil
}
