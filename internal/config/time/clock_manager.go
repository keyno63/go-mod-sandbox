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

type ClockManagerMock struct {
}

func (cm ClockManagerMock) Now() int64 {
	return int64(1)
}

type ClockUser struct {
	ClockM ClockManager
}
