package config

import (
	"encoding/json"
	"regexp"
	"time"
)

type Config struct {
	CurrentUserID      uint64    `json:"current_user_id"`
	CurrentUserSetTime time.Time `json:"current_user_set_time"`
	Users              []User    `json:"users"`
}

type User struct {
	ID       uint64   `json:"id"`
	APIToken string   `json:"api_token"`
	Name     string   `json:"name"`
	Username string   `json:"username"`
	Aliases  []string `json:"aliases,omitempty"`
}

func (config *Config) AddUser(id uint64, apiToken string, name string, username string, aliases []string) error {
	//TODO: Turn this into a flag helper
	valid, err := regexp.MatchString("^[a-fA-F\\d]{32}$", apiToken)
	if err != nil {
		return err
	}
	if !valid {
		return InvalidAPITokenError
	}

	user := User{
		ID:       id,
		APIToken: apiToken,
		Name:     name,
		Username: username,
		Aliases:  aliases,
	}

	for i, savedUser := range config.Users {
		if savedUser.ID == user.ID {
			config.Users = append(config.Users[:i], config.Users[i+1:]...)
			break
		}
	}

	for _, savedUser := range config.Users {
		for _, savedAlias := range savedUser.Aliases {
			for _, alias := range aliases {
				if savedAlias == alias {
					return DuplicateAliasError{User: savedUser}
				}
			}
		}
	}

	config.Users = append(config.Users, user)
	return nil
}

func (config *Config) SetCurrentUser(username string) error {
	if username == "" {
		config.CurrentUserID = 0
		config.CurrentUserSetTime = time.Time{}
		return nil
	}

	for _, user := range config.Users {
		if user.Username == username {
			config.CurrentUserID = user.ID
			config.CurrentUserSetTime = time.Now()
			return nil
		}
	}

	return UserDoesNotExistError
}

func LoadConfig(rawConfig []byte) (Config, error) {
	var config Config
	err := json.Unmarshal(rawConfig, &config)
	return config, err
}

func SaveConfig(config Config) ([]byte, error) {
	return json.Marshal(config)
}
