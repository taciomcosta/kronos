package domain

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/taciomcosta/kronos/internal/domain/ticker"
)

type Job struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Tick    string `json:"tick"`
	ticker  ticker.Ticker
}

func NewJob(request CreateJobRequest) (Job, error) {
	ticker, err := ticker.NewTicker(request.Tick)
	if err != nil {
		return Job{}, err
	}
	return Job{request.Name, request.Command, request.Tick, ticker}, nil
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

type CreateJobRequest struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Tick    string `json:"tick"`
}

func CreateJob(request CreateJobRequest) error {
	fmt.Printf("Creating job %s\n", request.Name)
	job, err := NewJob(request)
	if err != nil {
		return err
	}
	err = repository.CreateJob(&job)
	if err != nil {
		return err
	}
	runner.AddJob(job)
	return nil
}
