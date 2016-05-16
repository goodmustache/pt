package config

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("Specified user does not exist in config.")

// DuplicateAliasError is the error that should be returned when adding a user
// but the alias provided is the same as an existing pt user.
type DuplicateAliasError struct {
	// User is the pre-existing pt user
	User User
}

// Error displays the error message
func (e DuplicateAliasError) Error() string {
	return fmt.Sprintf("Alias provided is already in use by %s (alias: %s)", e.User.Name, e.User.Alias)
}

// UsernameMatchesSavedAliasError is the error that should be returned when
// adding a user but the username is the same as an existing user's alias.
type UsernameMatchesSavedAliasError struct {
	// SavedUser is the pre-existing pt user
	SavedUser User

	// NewUser is the user to be added to pt
	NewUser User
}

// Error displays the error message
func (e UsernameMatchesSavedAliasError) Error() string {
	return fmt.Sprintf(
		"The user %s (%s) cannot be added because it's username matches %s (alias: %s)",
		e.NewUser.Name,
		e.NewUser.Username,
		e.SavedUser.Name,
		e.SavedUser.Alias,
	)
}

// AliasMatchesSavedUsernameError is the error that should be returned when
// adding a user but the alias is the same as an existing user's username.
type AliasMatchesSavedUsernameError struct {
	// SavedUser is the pre-existing pt user
	SavedUser User

	// NewUser is the user to be added to pt
	NewUser User
}

// Error displays the error message
func (e AliasMatchesSavedUsernameError) Error() string {
	return fmt.Sprintf(
		"The user %s (alias: %s) cannot be added because it's alias matches %s (%s)",
		e.NewUser.Name,
		e.NewUser.Alias,
		e.SavedUser.Name,
		e.SavedUser.Username,
	)
}
