package command

type UserList struct {
	Config Config
	UI     UI
}

func (cmd UserList) Execute(_ []string) error {
	configuredUsers, err := cmd.Config.GetUsers()
	if err != nil {
		return err
	}

	type DisplayUser struct {
		Username string `header:"username"`
		Name     string `header:"name"`
		Email    string `header:"email"`
	}

	users := []DisplayUser{}
	for _, user := range configuredUsers {
		users = append(users, DisplayUser{
			Username: user.Username,
			Name:     user.Name,
			Email:    user.Email,
		})
	}

	cmd.UI.PrintTable(users)

	return nil
}
