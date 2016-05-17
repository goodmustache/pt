package tracker

import (
	"errors"
	"net/http"
	"regexp"
)

// APIToken is an API Key that represents the user
type APIToken string

// ErrorInvalidAPIToken is returned when the API token does not match the
// expected format
var ErrorInvalidAPIToken = errors.New("API Token must be a 32 character long hex string. (Example: '1234567890abcdef1234567890abcdef')")

// Validate will return an ErrorInvalidAPIToken if the token does not match the
// expected format. The expected format is a 32 character hexadecimal string.
func (token APIToken) Validate() error {
	if !regexp.MustCompile("^[a-fA-F\\d]{32}$").MatchString(string(token)) {
		return ErrorInvalidAPIToken
	}

	return nil
}

// AddToRequestHeader will add the APIToken to a given request with the
// appropriate header.
func (token APIToken) AddToRequestHeader(request *http.Request) {
	request.Header.Add("X-TrackerToken", string(token))
}
