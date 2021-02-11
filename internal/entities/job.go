package entities

import (
	"time"
)

// NewJob creates a new Job using options
func NewJob(name string, command string, tick string, stream Stream) (Job, error) {
	ticker, err := NewTicker(tick)
	job := Job{
		Name:    name,
		Command: command,
		Tick:    tick,
		ticker:  ticker,
		stream:  stream,
	}
	return job, err
}

// Job represents a job
type Job struct {
	Name    string
	Command string
	Tick    string
	ticker  Ticker
	stream  Stream
}

// IsTimeSet tells if job should run in time t
func (j *Job) IsTimeSet(t time.Time) bool {
	return j.ticker.IsTimeSet(t)
}
