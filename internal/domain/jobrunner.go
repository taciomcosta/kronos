package domain

import (
	"fmt"
	"time"
)

// JobsRunner acts as a container for managing many jobs at once.
type JobsRunner struct {
	jobs []Job
}

// NewJobRunner creates a new JobRunner.
func NewJobRunner(jobs []Job) JobsRunner {
	return JobsRunner{jobs}
}

// AddJob adds a Job to JobRunner, making it scheduled.
func (jr *JobsRunner) AddJob(job Job) {
	fmt.Println("new job added")
	jr.jobs = append(jr.jobs, job)
}

// Start starts runner on system startup.
func (jr *JobsRunner) Start() {
	jr.tickForever()
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
