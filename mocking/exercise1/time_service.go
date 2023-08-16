package exercise1

import "time"

// TimeProvider is an interface for interacting with time-related operations.
type TimeProvider interface {
	Now() time.Time
}

// RealTimeProvider is the implementation of TimeProvider that uses the real time.
type RealTimeProvider struct{}

// Now returns the current time.
func (rtp *RealTimeProvider) Now() time.Time {
	return time.Now()
}

func GetCurrentTime(tp TimeProvider) time.Time {
	return tp.Now()
}
