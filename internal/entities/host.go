package entities

import "time"

// Host represents a host where jobs can be run
type Host interface {
	RunJob(job Job)
	TickEverySecond() <-chan time.Time
}
