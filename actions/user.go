package actions

import (
	"errors"
	"os"

	"github.com/goodmustache/pt/config"
)

var ErrUnableToFindAlias = errors.New("Unable to find alias in config.")
var ErrUnableToFindUsername = errors.New("Unable to find username in config.")
var ErrNoCurrentUserSet = errors.New("There is no current user set in the config, either add a user (via 'add-user') or set the current user (via 'set-user').")
var ErrBothAliasAndUsernameProvided = errors.New("Both alias and username were provided, only one is allowed.")

type ConfigUser config.User

func AddUser(client TrackerClient, alias string) (ConfigUser, error) {
	tokenInfo, err := client.TokenInfo()
	if err != nil {
		return ConfigUser{}, err
	}

	conf, err := ReadConfig()
	if err != nil && !os.IsNotExist(err) {
		return ConfigUser{}, err
	}

	user := ConfigUser{
		ID:       tokenInfo.ID,
		APIToken: tokenInfo.APIToken,
		Name:     tokenInfo.Name,
		Username: tokenInfo.Username,
		Alias:    alias,
	}

	err = conf.AddUser(user.ID, user.APIToken, user.Name, user.Username, user.Alias)
	if err != nil {
		return ConfigUser{}, err
	}

	err = conf.SetCurrentUser(user.Username)
	if err != nil {
		return ConfigUser{}, err
	}

	err = WriteConfig(conf)
	if err != nil {
		return ConfigUser{}, err
	}

	return user, nil
}

func GetUser(alias string, username string) (ConfigUser, error) {
	conf, err := ReadConfig()
	if err != nil {
		if exists := os.IsNotExist(err); exists {
			return ConfigUser{}, ErrNoCurrentUserSet
		}
		return ConfigUser{}, err
	}

	if alias != "" && username != "" {
		return ConfigUser{}, ErrBothAliasAndUsernameProvided
	}

	switch {
	case alias != "":
		for _, user := range conf.Users {
			if user.Alias == alias {
				return ConfigUser(user), nil
			}
		}
		return ConfigUser{}, ErrUnableToFindAlias

	case username != "":
		for _, user := range conf.Users {
			if user.Username == username {
				return ConfigUser(user), nil
			}
		}
		return ConfigUser{}, ErrUnableToFindUsername

	default:
		for _, user := range conf.Users {
			if user.ID == conf.CurrentUserID {
				return ConfigUser(user), nil
			}
		}
		return ConfigUser{}, ErrNoCurrentUserSet
	}
}

func RemoveUser(userToRemove ConfigUser) error {
	conf, err := ReadConfig()
	if err != nil {
		return err
	}

	err = conf.RemoveUser(config.User(userToRemove))
	if err != nil {
		return err
	}

	return WriteConfig(conf)
}
