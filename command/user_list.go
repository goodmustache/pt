package command

import "github.com/goodmustache/pt/command/display"

type UserList struct {
	Config Config
	UI     UI
}

func (cmd UserList) Execute(_ []string) error {
	configuredUsers, err := cmd.Config.GetUsers()
	if err != nil {
		return err
	}

	users := []display.UserRow{}
	for _, user := range configuredUsers {
		users = append(users, display.UserRow{
			ID:       user.ID,
			Username: user.Username,
			Name:     user.Name,
			Email:    user.Email,
		})
	}

	cmd.UI.PrintTable(users)

	return nil
}
