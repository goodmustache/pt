package tracker

import (
	"time"

	"github.com/goodmustache/pt/tracker/internal"
)

type Me struct {
	// APIToken is a string that can be used as the API authentication token
	// (X-TrackerToken) to authenticate future API requests as being on behalf of
	// the current user. This field is read only.
	APIToken string `json:"api_token"`
	// Email is the authenticated user's email. This field is read only.
	Email string `json:"email"`
	// ID is the database id of the authenticated user. This field is read only.
	// This field is always returned.
	ID uint64 `json:"id"`
	// Name is the name of the authenticated user. This field is read only.
	Name string `json:"name"`
	// Projects is a list of the project(s) that the authenticated user is a
	// member of. This field is read only.
	Projects []EmbeddedProject `json:"projects"`
	// Username is the authenticated user's optional 'username' for login
	// purposes. This field is read only.
	Username string `json:"username"`
}

type EmbeddedProject struct {
	Color        string    `json:"color"`
	Favorite     bool      `json:"favorite"`
	ID           uint64    `json:"id"`
	Kind         string    `json:"kind"`
	LastViewedAt time.Time `json:"last_viewed_at"`
	Name         string    `json:"project_name"`
	ProjectID    uint64    `json:"project_id"`
	Role         string    `json:"role"`
}

// Me returns information from the user's profile plus the list of projects
// to which the user has access.
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
