package actions

import "github.com/goodmustache/pt/tracker"

//go:generate counterfeiter . TrackerClient

// TrackerClient is the API client used to talk to Pivotal Tracker
type TrackerClient interface {
	TokenInformation() (tracker.TokenInformation, error)
}
