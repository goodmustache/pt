package internal

import "time"

//counterfeiter:generate . Clock

type Clock interface {
	Now() time.Time
}
