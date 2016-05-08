package config

import (
	"encoding/json"
	"time"
)

type Config struct {
	CurrentUserID      uint64    `json:"current_user_id"`
	CurrentUserSetTime time.Time `json:"current_user_set_time"`
	Users              []User    `json:"users"`
}

type User struct {
	ID       uint64 `json:"id"`
	APIToken string `json:"api_token"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Alias    string `json:"alias,omitempty"`
}

func (config *Config) AddUser(id uint64, apiToken string, name string, username string, alias string) error {
	newUser := User{
		ID:       id,
		APIToken: apiToken,
		Name:     name,
		Username: username,
		Alias:    alias,
	}

	for i, savedUser := range config.Users {
		if savedUser.ID == newUser.ID {
			config.Users = append(config.Users[:i], config.Users[i+1:]...)
			break
		}
	}

	for _, savedUser := range config.Users {
		switch {
		case newUser.Alias != "" && savedUser.Alias == newUser.Alias:
			return DuplicateAliasError{User: savedUser}
		case savedUser.Alias == newUser.Username:
			return UsernameMatchesSavedAliasError{SavedUser: savedUser, NewUser: newUser}
		case savedUser.Username == newUser.Alias:
			return AliasMatchesSavedUsernameError{SavedUser: savedUser, NewUser: newUser}
		}
	}

	config.Users = append(config.Users, newUser)
	return nil
}

func (config *Config) RemoveUser(userToRemove User) error {
	var found bool
	for i, user := range config.Users {
		if user.ID == userToRemove.ID {
			config.Users = append(config.Users[:i], config.Users[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return ErrUserNotFound
	}

	if config.CurrentUserID == userToRemove.ID {
		return config.SetCurrentUser("")
	}
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

	return ErrUserNotFound
}

func LoadConfig(rawConfig []byte) (Config, error) {
	var config Config
	err := json.Unmarshal(rawConfig, &config)
	return config, err
}

func SaveConfig(config Config) ([]byte, error) {
	return json.Marshal(config)
}
