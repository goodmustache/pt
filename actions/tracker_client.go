package actions

import "github.com/goodmustache/pt/tracker"

//go:generate counterfieter . TrackerClient

type TrackerClient interface {
	TokenInfo() (tracker.TokenInfomation, error)
}
