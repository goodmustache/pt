package actor

import "github.com/goodmustache/pt/tracker"

//counterfeiter:generate . TrackerClient

type TrackerClient interface {
	Me() (tracker.Me, error)
	Projects() ([]tracker.Project, error)
}
