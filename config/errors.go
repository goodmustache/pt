package config

import (
	"errors"
	"fmt"
)

var UserDoesNotExistError = errors.New("Specified user does not exist in config.")

type DuplicateAliasError struct {
	User User
}

func (e DuplicateAliasError) Error() string {
	return fmt.Sprintf("Alias provided is already in use by %s (alias: %s)", e.User.Name, e.User.Alias)
}
