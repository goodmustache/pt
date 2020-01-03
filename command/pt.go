package command

type CommandList struct {
	User struct {
		List UserList `command:"list" description:"list all known users"`
	} `command:"user" description:"local user related commands"`
}

var PT CommandList
