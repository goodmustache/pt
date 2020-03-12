package tracker

import "github.com/goodmustache/pt/tracker/internal"

type Me struct {
	APIToken string `json:"api_token"`
	Email    string `json:"email"`
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func (client *Client) Me() (Me, error) {
	var me Me
	return me,
		client.MakeRequest(
			requestOptions{
				RequestName: internal.GetMeRequest,
			},
			&me,
		)
}
