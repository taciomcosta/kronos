package domain

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type JobsRunner struct {
	jobs []Job
}

func NewJobRunner() JobsRunner {
	return JobsRunner{}
}

func (jr *JobsRunner) Start() {
	ticker := time.NewTicker(1 * time.Second)
	last := time.Now().UTC()
	for range ticker.C {
		now := time.Now().UTC()
		if now.Minute() != last.Minute() {
			last = now
			jr.runAllJobs(now)
		}
	}
}

func (jr *JobsRunner) runAllJobs(t time.Time) {
	fmt.Printf("Starting execution at %v\n", t)
	for _, job := range jr.jobs {
		go job.Run(t)
	}
}

func (jr *JobsRunner) AddJob(job Job) {
	jr.jobs = append(jr.jobs, job)
}

type Job struct {
	Name    string
	Command string
	Tick    string
	ticker  Ticker
}

func NewJob(name string, command string, tick string) (Job, error) {
	ticker, err := NewTicker(tick)
	if err != nil {
		return Job{}, err
	}
	return Job{name, command, tick, ticker}, nil
}

func (j *Job) Run(t time.Time) {
	if !j.ticker.IsTimeSet(t) {
		return
	}
	fmt.Printf("Executing %s\n", j.Name)
	j.ExecCommand()
}

func (j *Job) ExecCommand() error {
	cmd := exec.Command(j.Command)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
