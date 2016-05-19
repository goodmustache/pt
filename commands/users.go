package commands

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/goodmustache/pt/actions"
)

// ErrNoUsers is returned when no users have been added to pt
var ErrNoUsers = errors.New("No users have been added, run the 'add-users' command.")

// UsersCommand lists the users in pt
type UsersCommand struct{}

// Execute is the execution of the UsersCommand
func (cmd UsersCommand) Execute([]string) error {
	config, err := actions.ReadConfig()
	if err != nil {
		if os.IsNotExist(err) {
			return ErrNoUsers
		}

		return err
	} else if len(config.Users) == 0 {
		return ErrNoUsers
	}

	var minimumColumnWidth int
	for _, user := range config.Users {
		if l := len(user.Name); l > minimumColumnWidth {
			minimumColumnWidth = l
		}

		if l := len(user.Username); l > minimumColumnWidth {
			minimumColumnWidth = l
		}
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, minimumColumnWidth, 8, 2, '\t', 0)
	fmt.Fprintln(w, "Name\tUsername\tAlias")

	for _, user := range config.Users {
		fmt.Fprintf(w, "%s\t%s\t%s\n", user.Name, user.Username, user.Alias)
	}

	w.Flush()

	return nil
}
