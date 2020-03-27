package command

type CommandList struct {
	Project struct {
		List ProjectList `command:"list" description:"list all accessable projects"`
	} `command:"project" description:"project related commands"`
	User struct {
		List UserList `command:"list" description:"list all known users"`
		Add  UserAdd  `command:"add" description:"add user to local machine"`
	} `command:"user" description:"local user related commands"`
}

var PT CommandList
