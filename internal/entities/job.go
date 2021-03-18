package entities

import (
	"time"
)

// NewJob creates a new Job using options
func NewJob(name string, command string, tick string) (Job, error) {
	ticker, err := NewTicker(tick)
	job := Job{
		Name:    name,
		Command: command,
		Tick:    tick,
		Status:  true,
		ticker:  ticker,
	}
	return job, err
}

// Job represents a job
type Job struct {
	Name    string
	Command string
	Tick    string
	Status  bool
	ticker  Ticker
}

// IsTimeSet tells if job should run in time t
func (j *Job) IsTimeSet(t time.Time) bool {
	return j.ticker.IsTimeSet(t)
}
