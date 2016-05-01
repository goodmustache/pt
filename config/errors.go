package config

import (
	"errors"
	"fmt"
	"strings"
)

var InvalidAPITokenError = errors.New("API Token must be a 32 character long hex string. (Example: '1234567890abcdef1234567890abcdef')")

var UserDoesNotExistError = errors.New("Specified user does not exist in config.")

type DuplicateAliasError struct {
	User User
}

func (e DuplicateAliasError) Error() string {
	return fmt.Sprintf("Alias provided is already in use by %s (aliases: %s)", e.User.Name, strings.Join(e.User.Aliases, ", "))
}
