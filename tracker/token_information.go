package tracker

import "encoding/json"

// TokenInformation matches a subset of the user information returned from the
// '/me' endpoint. For more information, see the Pivotal Tracker API Docs:
// https://www.pivotaltracker.com/help/api/rest/v5#Me
type TokenInformation struct {
	// ID is the user's assigned ID
	ID uint64 `json:"id"`

	// APIToken is the user's assigned API Token
	APIToken string `json:"api_token"`

	// Name is the user provided name
	Name string `json:"name"`

	// Username is the user provided username
	Username string `json:"username"`
}

// TokenInformation GETs the '/me' endpoint and returns a subset of the
// information provided.
func (c Client) TokenInformation() (TokenInformation, error) {
	responseBody, err := c.get("me")
	if err != nil {
		return TokenInformation{}, err
	}

	var tokenInfo TokenInformation
	err = json.Unmarshal(responseBody, &tokenInfo)
	return tokenInfo, err
}
