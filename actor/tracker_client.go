package actor

import "github.com/goodmustache/pt/tracker"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . TrackerClient

type TrackerClient interface {
	Me() (tracker.Me, error)
}
