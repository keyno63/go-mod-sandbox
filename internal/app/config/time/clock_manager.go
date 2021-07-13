package time

import "time"

type ClockManager interface {
	Now() int64
}

type ClockManagerImpl struct {
}

func (cm ClockManagerImpl) Now() int64 {
	return time.Now().Unix()
}
