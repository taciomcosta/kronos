package domain

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/taciomcosta/kronos/internal/domain/ticker"
)

func NewJob(request CreateJobRequest) (Job, error) {
	ticker, err := ticker.NewTicker(request.Tick)
	if err != nil {
		return Job{}, err
	}
	return Job{request.Name, request.Command, request.Tick, ticker}, nil
}

type CreateJobRequest struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Tick    string `json:"tick"`
}

type CreateJobResponse struct {
	Msg string `json:"msg"`
}

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

func FindJobs() []Job {
	return repository.FindJobs()
}

func CountJobs() int {
	return repository.CountJobs()
}

type Job struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Tick    string `json:"tick"`
	ticker  ticker.Ticker
}

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
