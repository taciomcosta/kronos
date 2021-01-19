package domain

import (
	"fmt"
	"time"
)

type JobsRunner struct {
	jobs []Job
}

func NewJobRunner() JobsRunner {
	return JobsRunner{}
}

func (jr *JobsRunner) AddJob(job Job) {
	jr.jobs = append(jr.jobs, job)
}

func (jr *JobsRunner) Start() {
	jr.loadJobs()
	jr.tickForever()
}

func (jr *JobsRunner) loadJobs() {
	for _, job := range repository.FindJobs() {
		jr.AddJob(job)
	}
}

func (jr *JobsRunner) tickForever() {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		now := time.Now().UTC()
		if now.Second() == 0 {
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
