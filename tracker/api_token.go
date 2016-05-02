package tracker

import (
	"errors"
	"net/http"
	"regexp"
)

type APIToken string

var ErrorInvalidAPIToken = errors.New("API Token must be a 32 character long hex string. (Example: '1234567890abcdef1234567890abcdef')")

func (token APIToken) Validate() error {
	if !regexp.MustCompile("^[a-fA-F\\d]{32}$").MatchString(string(token)) {
		return ErrorInvalidAPIToken
	}

	return nil
}

func (token APIToken) AddToRequestHeader(request *http.Request) {
	request.Header.Add("X-TrackerToken", string(token))
}
