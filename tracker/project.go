package tracker

import "github.com/goodmustache/pt/tracker/internal"

type Project struct {
	Description string `json:"description"`
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Public      bool   `json:"public"`
}

func (client *Client) Projects() ([]Project, error) {
	var projects []Project
	return projects,
		client.MakeRequest(
			requestOptions{
				RequestName: internal.GetProjectsRequest,
			},
			&projects,
		)
}
