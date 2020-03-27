package command

import (
	"github.com/goodmustache/pt/actor"
	"github.com/goodmustache/pt/command/display"
)

//counterfeiter:generate . ProjectListActor

type ProjectListActor interface {
	Projects() ([]actor.Project, error)
}

type ProjectList struct {
	UserID uint64 `short:"u" long:"user-id" description:"User ID to run commands with"`

	Actor ProjectListActor
	UI    UI
}

func (cmd ProjectList) Execute(_ []string) error {
	projects, err := cmd.Actor.Projects()
	if err != nil {
		return err
	}

	displayProjects := make([]display.ProjectRow, 0, len(projects))
	for _, project := range projects {
		displayProjects = append(displayProjects, display.ProjectRow{
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description,
			Visibility:  project.Visibility(),
		})
	}

	cmd.UI.PrintTable(displayProjects)
	return nil
}
