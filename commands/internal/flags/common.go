package flags

// CommonFlags are common flags used across many commands
type CommonFlags struct {
	Alias    string `short:"a" long:"alias" description:"Run command as user with given alias"`
	Username string `short:"u" long:"username" description:"Run command as user with given username"`
}
