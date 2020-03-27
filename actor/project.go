package actor

import (
	"sort"
	"strings"

	"github.com/goodmustache/pt/tracker"
)

type Project tracker.Project

func (project Project) Visibility() string {
	if project.Public {
		return "public"
	}
	return "private"
}

func (actor *Main) Projects() ([]Project, error) {
	apiProjects, err := actor.API.Projects()
	projects := make([]Project, 0, len(apiProjects))
	for _, p := range apiProjects {
		projects = append(projects, Project(p))
	}

	sort.Slice(projects,
		func(i, j int) bool {
			return strings.ToLower(projects[i].Name) < strings.ToLower(projects[j].Name)
		})

	return projects, err
}
