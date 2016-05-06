package internal

import (
	"errors"
	"os"

	"github.com/goodmustache/pt/config"
)

var ErrUnableToFindAlias = errors.New("Unable to find alias in config.")
var ErrUnableToFindUsername = errors.New("Unable to find username in config.")
var ErrNoCurrentUserSet = errors.New("There is no current user set in the config, either add a user (via 'add-user') or set the current user (via 'set-user').")
var ErrBothAliasAndUsernameProvided = errors.New("Both alias and username were provided, only one is allowed.")

func GetUser(alias string, username string) (config.User, error) {
	conf, err := config.ReadConfig()
	if err != nil {
		if exists := os.IsNotExist(err); exists {
			return config.User{}, ErrNoCurrentUserSet
		}
		return config.User{}, err
	}

	if alias != "" && username != "" {
		return config.User{}, ErrBothAliasAndUsernameProvided
	}

	switch {
	case alias != "":
		for _, user := range conf.Users {
			if user.Alias == alias {
				return user, nil
			}
		}
		return config.User{}, ErrUnableToFindAlias

	case username != "":
		for _, user := range conf.Users {
			if user.Username == username {
				return user, nil
			}
		}
		return config.User{}, ErrUnableToFindUsername

	default:
		for _, user := range conf.Users {
			if user.ID == conf.CurrentUserID {
				return user, nil
			}
		}
		return config.User{}, ErrNoCurrentUserSet
	}
}
