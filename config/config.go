package config

import (
	"encoding/json"
	"time"
)

// Config holds metadata about pt's users
type Config struct {
	// CurrentUserID is the current user's ID
	CurrentUserID uint64 `json:"current_user_id"`

	// CurrentUserSetTime is when the current user is set
	CurrentUserSetTime time.Time `json:"current_user_set_time"`

	// Users is the list of pt's users
	Users []User `json:"users"`
}

// User is a PT system user
type User struct {
	// ID given from Tracker
	ID uint64 `json:"id"`

	// APIToken provided from the Tracker's settings page
	APIToken string `json:"api_token"`

	// Name is the real name of the user
	Name string `json:"name"`

	// Username for Tracker
	Username string `json:"username"`

	// Alias setup by the user for quick reference
	Alias string `json:"alias,omitempty"`
}

// AddUser adds users to config with the given credentials. Will error out if:
//	- API Token matches another user's API token (DuplicateAliasError)
// 	- Username matches a saved alias (UsernameMatchesSavedAliasError)
// 	- Alias matches another user's alias (AliasMatchesSavedUsernameError)
// If a user is added again (determined by the same ID), the original user will
// be removed from the config.
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

// RemoveUser removes the given user from the config if it's ID matches one in
// the config.
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

// SetCurrentUser sets the current user on the config.
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

// LoadConfig converts a config from a raw byte array.
func LoadConfig(rawConfig []byte) (Config, error) {
	var config Config
	err := json.Unmarshal(rawConfig, &config)
	return config, err
}

// SaveConfig converts a raw config into a byte array.
func SaveConfig(config Config) ([]byte, error) {
	return json.Marshal(config)
}
