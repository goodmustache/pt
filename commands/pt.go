package commands

type PTCommand struct {
	TrackerURL string `default:"https://www.pivotaltracker.com/services/v5" long:"override-tracker-url" hidden:"true"`

	AddUser AddUserCommand `command:"add-user" alias:"au" description:"Add user's API key and login"`
}

var PT PTCommand
