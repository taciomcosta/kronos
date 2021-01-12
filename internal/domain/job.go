package domain

import (
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
	Msg     string `json:"msg"`
	Success bool
}

func CreateJob(request CreateJobRequest) CreateJobResponse {
	job, err := NewJob(request)
	if err != nil {
		return CreateJobResponse{Msg: err.Error(), Success: false}
	}
	err = repository.CreateJob(&job)
	if err != nil {
		return CreateJobResponse{}
	}
	runner.AddJob(job)
	return CreateJobResponse{Msg: job.Name + " created.", Success: true}
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
