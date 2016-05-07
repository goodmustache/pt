package actions

import (
	"errors"
	"os"
	"time"

	"github.com/goodmustache/pt/config"
)

var ErrUnableToFindAlias = errors.New("Unable to find alias in config.")
var ErrUnableToFindUsername = errors.New("Unable to find username in config.")
var ErrNoCurrentUserSet = errors.New("There is no current user set in the config, either add a user (via 'add-user') or set the current user (via 'set-user').")
var ErrBothAliasAndUsernameProvided = errors.New("Both alias and username were provided, only one is allowed.")

type User struct {
	ID       uint64 `json:"id"`
	APIToken string `json:"api_token"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Alias    string `json:"alias,omitempty"`
}

func AddUser(client TrackerClient, alias string) (User, error) {
	tokenInfo, err := client.TokenInfo()
	if err != nil {
		return User{}, err
	}

	conf, err := config.ReadConfig()
	if err != nil && !os.IsNotExist(err) {
		return User{}, err
	}

	user := User{
		ID:       tokenInfo.ID,
		APIToken: tokenInfo.APIToken,
		Name:     tokenInfo.Name,
		Username: tokenInfo.Username,
		Alias:    alias,
	}

	err = conf.AddUser(user.ID, user.APIToken, user.Name, user.Username, user.Alias)
	if err != nil {
		return User{}, err
	}

	err = conf.SetCurrentUser(user.Username)
	if err != nil {
		return User{}, err
	}

	err = config.WriteConfig(conf)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func GetUser(alias string, username string) (User, error) {
	conf, err := config.ReadConfig()
	if err != nil {
		if exists := os.IsNotExist(err); exists {
			return User{}, ErrNoCurrentUserSet
		}
		return User{}, err
	}

	if alias != "" && username != "" {
		return User{}, ErrBothAliasAndUsernameProvided
	}

	switch {
	case alias != "":
		for _, user := range conf.Users {
			if user.Alias == alias {
				return User(user), nil
			}
		}
		return User{}, ErrUnableToFindAlias

	case username != "":
		for _, user := range conf.Users {
			if user.Username == username {
				return User(user), nil
			}
		}
		return User{}, ErrUnableToFindUsername

	default:
		for _, user := range conf.Users {
			if user.ID == conf.CurrentUserID {
				return User(user), nil
			}
		}
		return User{}, ErrNoCurrentUserSet
	}
}

func RemoveUser(userToRemove User) error {
	conf, err := config.ReadConfig()
	if err != nil {
		return err
	}

	for i, user := range conf.Users {
		if user.ID == userToRemove.ID {
			conf.Users = append(conf.Users[:i], conf.Users[i+1:]...)
			break
		}
	}

	if conf.CurrentUserID == userToRemove.ID {
		conf.CurrentUserID = 0
		conf.CurrentUserSetTime = time.Time{}
	}

	return config.WriteConfig(conf)
}
