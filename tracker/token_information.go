package tracker

import "encoding/json"

type TokenInfomation struct {
	ID       uint64 `json:"id"`
	APIToken string `json:"api_token"`
	Initials string `json:"initials"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func (c client) TokenInfo() (TokenInfomation, error) {
	var tokenInfo TokenInfomation
	responseBody, err := c.get("/me")
	if err != nil {
		return tokenInfo, err
	}

	err = json.Unmarshal(responseBody, &tokenInfo)
	return tokenInfo, err
}
