package entities

import (
	"time"
)

// NewJob creates a new Job using options
func NewJob(name string, command string, tick string, status bool) (Job, error) {
	ticker, err := NewTicker(tick)
	job := Job{
		Name:    name,
		Command: command,
		Tick:    tick,
		Status:  status,
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

// ShouldRun checks provided time and job status to see if it should run
func (j *Job) ShouldRun(t time.Time) bool {
	return j.isEnabled() && j.ticker.IsTimeSet(t)
}

func (j *Job) isEnabled() bool {
	return j.Status
}
