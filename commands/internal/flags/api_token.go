package flags

import "github.com/goodmustache/pt/tracker"

// APIToken is an input validator and parser for a Pivotal Tracker API token
type APIToken struct {
	Value tracker.APIToken
}

// UnmarshalFlag will parse then validate a token value
func (token *APIToken) UnmarshalFlag(value string) error {
	token.Value = tracker.APIToken(value)
	return token.Value.Validate()
}
