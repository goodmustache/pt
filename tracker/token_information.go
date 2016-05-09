package tracker

import "encoding/json"

type TokenInformation struct {
	ID       uint64 `json:"id"`
	APIToken string `json:"api_token"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func (c Client) TokenInformation() (TokenInformation, error) {
	responseBody, err := c.get("/me")
	if err != nil {
		return TokenInformation{}, err
	}

	var tokenInfo TokenInformation
	err = json.Unmarshal(responseBody, &tokenInfo)
	return tokenInfo, err
}
