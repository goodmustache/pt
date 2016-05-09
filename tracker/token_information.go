package tracker

import "encoding/json"

type TokenInformation struct {
	ID       uint64 `json:"id"`
	APIToken string `json:"api_token"`
	Initials string `json:"initials"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func (c client) TokenInformation() (TokenInformation, error) {
	responseBody, err := c.get("/me")
	if err != nil {
		return TokenInformation{}, err
	}

	var tokenInfo TokenInformation
	err = json.Unmarshal(responseBody, &tokenInfo)
	return tokenInfo, err
}
