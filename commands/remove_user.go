package commands

import (
	"fmt"
	"time"

	"github.com/goodmustache/pt/commands/internal"
	"github.com/goodmustache/pt/config"
	"github.com/vito/go-interact/interact"
)

type RemoveUserCommand struct {
	Alias    string `short:"a" long:"alias" description:"Target user's alias"`
	Username string `short:"u" long:"username" description:"Target user's username"`
	Force    bool   `long:"with-malice" description:"Remove target user without confirmation"`
}

func (cmd *RemoveUserCommand) Execute([]string) error {
	selectedUser, err := internal.GetUser(cmd.Alias, cmd.Username)
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

	conf, err := config.ReadConfig()

	for i, user := range conf.Users {
		if user.ID == selectedUser.ID {
			conf.Users = append(conf.Users[:i], conf.Users[i+1:]...)
		}
	}

	if conf.CurrentUserID == selectedUser.ID {
		conf.CurrentUserID = 0
		conf.CurrentUserSetTime = time.Time{}
	}

	err = config.WriteConfig(conf)
	if err != nil {
		return err
	}

	fmt.Printf("User %s (%s) has been removed.\n", selectedUser.Name, selectedUser.Username)

	return nil
}
