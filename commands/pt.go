package commands

type PTCommand struct {
	TrackerURL string `default:"https://www.pivotaltracker.com/services/v5/" long:"override-tracker-url" hidden:"true"`

	AddUser    AddUserCommand    `command:"add-user" alias:"au" description:"Add user's API token and login"`
	RemoveUser RemoveUserCommand `command:"remove-user" alias:"ru" description:"Remove user's API token"`
	User       UserCommand       `command:"user" alias:"u" description:"Displays pt user"`
	Users      UsersCommand      `command:"users" description:"List all pt users"`
}

var PT PTCommand
