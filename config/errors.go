package config

import (
	"errors"
	"fmt"
)

var ErrorUserDoesNotExist = errors.New("Specified user does not exist in config.")

type DuplicateAliasError struct {
	User User
}

func (e DuplicateAliasError) Error() string {
	return fmt.Sprintf("Alias provided is already in use by %s (alias: %s)", e.User.Name, e.User.Alias)
}

type UsernameMatchesSavedAliasError struct {
	SavedUser User
	NewUser   User
}

func (e UsernameMatchesSavedAliasError) Error() string {
	return fmt.Sprintf(
		"The user %s (%s) cannot be added because it's username matches %s (alias: %s)",
		e.NewUser.Name,
		e.NewUser.Username,
		e.SavedUser.Name,
		e.SavedUser.Alias,
	)
}

type AliasMatchesSavedUsernameError struct {
	SavedUser User
	NewUser   User
}

func (e AliasMatchesSavedUsernameError) Error() string {
	return fmt.Sprintf(
		"The user %s (alias: %s) cannot be added because it's alias matches %s (%s)",
		e.NewUser.Name,
		e.NewUser.Alias,
		e.SavedUser.Name,
		e.SavedUser.Username,
	)
}
