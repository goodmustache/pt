package actions

import "github.com/goodmustache/pt/tracker"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . TrackerClient

// TrackerClient is the API client used to talk to Pivotal Tracker
type TrackerClient interface {
	TokenInformation() (tracker.TokenInformation, error)
}
