package commands

import (
	"fmt"
	"os"

	"github.com/goodmustache/pt/commands/internal"
	"github.com/goodmustache/pt/commands/internal/flags"
	"github.com/goodmustache/pt/tracker"
	"github.com/vito/go-interact/interact"
)

const AddUserInstructions = `In order to add a user, you must provide an API Token. This can be found at the bottom of your Profile page, under the "API TOKEN" section. If none is listed, click the "CREATE NEW TOKEN" button. To find the Profile page, follow the follwing link:

	https://www.pivotaltracker.com/profile
`

type AddUserCommand struct {
	APIToken flags.APIToken `long:"api-token" describe:"API Token for a user"`
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

	client := tracker.NewTrackerClient(PT.TrackerURL, apiToken)
	tokenInfo, err := client.TokenInfo()
	if err != nil {
		return err
	}

	config, err := internal.ReadConfig()
	if exists := os.IsNotExist(err); !exists {
		return err
	}

	err = config.AddUser(tokenInfo.ID, tokenInfo.APIToken, tokenInfo.Name, tokenInfo.Username, []string{})
	if err != nil {
		return err
	}

	err = config.SetCurrentUser(tokenInfo.Username)
	if err != nil {
		return err
	}

	err = internal.WriteConfig(config)
	if err != nil {
		return err
	}

	fmt.Printf("Added User! Setting %s (%s) to be the current user.\n", tokenInfo.Name, tokenInfo.Username)

	return nil
}
