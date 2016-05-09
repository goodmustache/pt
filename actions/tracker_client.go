package actions

import "github.com/goodmustache/pt/tracker"

//go:generate counterfeiter . TrackerClient

type TrackerClient interface {
	TokenInfo() (tracker.TokenInfomation, error)
}
