package flags

import "github.com/goodmustache/pt/tracker"

type APIToken struct {
	Value tracker.APIToken
}

func (token *APIToken) UnmarshalFlag(value string) error {
	token.Value = tracker.APIToken(value)
	return token.Value.Validate()
}
