package command

import (
	"github.com/goodmustache/pt/actor"
	"github.com/goodmustache/pt/command/display"
)

//counterfeiter:generate . UserAddActor

type UserAddActor interface {
	AddUser() (actor.User, error)
}

type UserAdd struct {
	APIToken string `long:"token" required:"true" description:"The Traker API Token for the user"`

	Actor  UserAddActor
	Config Config
	UI     UI
}

func (cmd UserAdd) Execute(_ []string) error {
	user, err := cmd.Actor.AddUser()
	if err != nil {
		return err
	}

	cmd.UI.PrintTable(display.UserRow{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	})

	return nil
}
