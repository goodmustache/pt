package commands

import (
	"fmt"

	"github.com/goodmustache/pt/actions"
	"github.com/goodmustache/pt/commands/internal/flags"
	"github.com/goodmustache/pt/tracker"
	"github.com/vito/go-interact/interact"
)

const AddUserInstructions = `In order to add a user, you must provide an API Token. This can be found at the bottom of your Profile page, under the "API TOKEN" section. If none is listed, click the "CREATE NEW TOKEN" button. To find the Profile page, follow the follwing link:

	https://www.pivotaltracker.com/profile
`

type AddUserCommand struct {
	APIToken flags.APIToken `long:"api-token" description:"API Token for a user"`
	Alias    string         `short:"a" long:"alias" description:"Alias to assign user"`
}

func (cmd *AddUserCommand) Execute([]string) error {
	apiToken := cmd.APIToken.Value

	if cmd.APIToken.Value == "" {
		fmt.Print(AddUserInstructions)

		var input string
		err := interact.NewInteraction("API Token").Resolve(interact.Required(&input))
		if err != nil {
			return err
		}

		apiToken = tracker.APIToken(input)
		err = apiToken.Validate()
		if err != nil {
			return err
		}
	}

	client := tracker.NewClient(PT.TrackerURL, apiToken)
	user, err := actions.AddUser(client, cmd.Alias)
	if err != nil {
		return err
	}

	fmt.Printf("Added User! Setting %s (%s) to be the current user.\n", user.Name, user.Username)

	return nil
}
