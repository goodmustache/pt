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
	responseBody, err := c.get("/me")
	if err != nil {
		return TokenInfomation{}, err
	}

	var tokenInfo TokenInfomation
	err = json.Unmarshal(responseBody, &tokenInfo)
	return tokenInfo, err
}
