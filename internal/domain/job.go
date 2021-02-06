package domain

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// NewJob creates a new Job from a request.
func NewJob(request CreateJobRequest) (Job, error) {
	ticker, err := NewTicker(request.Tick)
	if err != nil {
		return Job{}, err
	}
	return Job{request.Name, request.Command, request.Tick, ticker}, nil
}

// CreateJobRequest represents the needed properties to create a Job
type CreateJobRequest struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Tick    string `json:"tick"`
}

// CreateJobResponse represents the response message of CreateJob
type CreateJobResponse struct {
	Msg string `json:"msg"`
}

// CreateJob creates a job and schedules it right away.
func CreateJob(request CreateJobRequest) (CreateJobResponse, error) {
	fmt.Printf("Count job is %v\n", CountJobs())
	job, err := NewJob(request)
	if err != nil {
		return CreateJobResponse{Msg: err.Error()}, err
	}
	err = repository.CreateJob(&job)
	if err != nil {
		return CreateJobResponse{Msg: err.Error()}, err
	}
	runner.AddJob(job)
	return CreateJobResponse{Msg: job.Name + " created."}, nil
}

// FindJobs returns a list of all jobs.
func FindJobs() []Job {
	return repository.FindJobs()
}

// CountJobs counts the total of jobs.
func CountJobs() int {
	return repository.CountJobs()
}

// Job represents a job
type Job struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Tick    string `json:"'btick"`
	ticker  Ticker
}

// Run runs a job if it is the appropriate time.
func (j *Job) Run(t time.Time) {
	if !j.ticker.IsTimeSet(t) {
		return
	}
	j.execCommand()
}

func (j *Job) execCommand() error {
	cmd := exec.Command(j.Command)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
