package actions

import "github.com/goodmustache/pt/tracker"

//go:generate counterfeiter . TrackerClient

type TrackerClient interface {
	TokenInformation() (tracker.TokenInformation, error)
}
