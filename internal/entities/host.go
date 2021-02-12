package entities

import "time"

// Host represents a host where jobs can be run
type Host interface {
	RunJob(job *Job)
	GetDettachedStream() Stream
	TickEverySecond() <-chan time.Time
}
