package commands

import "errors"

// PTCommand is the root of pt, it lists all the available commands and some
// global settings
type PTCommand struct {
	TrackerURL string `default:"https://www.pivotaltracker.com/services/v5/" long:"override-tracker-url" hidden:"true"`

	AddUser    AddUserCommand    `command:"add-user" description:"Add user's API token and login"`
	RemoveUser RemoveUserCommand `command:"remove-user" description:"Remove user's API token"`
	User       UserCommand       `command:"user" description:"Displays pt user"`
	Users      UsersCommand      `command:"users" description:"List all pt users"`
}

// ErrUserTerminated is returned when a user cancels certain operations
var ErrUserTerminated = errors.New("User Terminated")

// PT is an instance of PTCommand that can be used to get global settings
var PT PTCommand
