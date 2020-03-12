package internal

import "time"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Clock

type Clock interface {
	Now() time.Time
}
