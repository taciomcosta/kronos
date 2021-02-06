package entities

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// NewJob creates a new Job from a request.
func NewJob(name, command, tick string) (Job, error) {
	ticker, err := NewTicker(tick)
	if err != nil {
		return Job{}, err
	}
	return Job{name, command, tick, ticker}, nil
}

// Job represents a job
type Job struct {
	Name    string
	Command string
	Tick    string
	ticker  Ticker
}

// Run runs a job if it is the appropriate time.
func (j *Job) Run(t time.Time) {
	if !j.ticker.IsTimeSet(t) {
		return
	}
	fmt.Printf("> Running %s\n", j.Name)
	j.execCommand()
}

func (j *Job) execCommand() error {
	cmd := exec.Command(j.Command)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
