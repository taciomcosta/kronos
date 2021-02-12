package usecases

import (
	"time"

	"github.com/taciomcosta/kronos/internal/entities"
)

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
	job, err := entities.NewJob(
		request.Name,
		request.Command,
		request.Tick,
		host.GetDettachedStream())
	if err != nil {
		return CreateJobResponse{}, err
	}
	err = repository.CreateJob(&job)
	if err != nil {
		return CreateJobResponse{}, err
	}
	jobs = append(jobs, job)
	return CreateJobResponse{Msg: job.Name + " created."}, nil
}

// FindJobs returns a list of all jobs.
func FindJobs() []entities.Job {
	return repository.FindJobs()
}

// CountJobs counts the total of jobs.
func CountJobs() int {
	return repository.CountJobs()
}

// ScheduleExistingJobs schedules jobs on startup
func ScheduleExistingJobs() {
	jobs = repository.FindJobs()
	go tickForever()
}

func tickForever() {
	for now := range host.TickEverySecond() {
		if now.Second() == 0 {
			runAllJobs(now)
		}
	}
}

func runAllJobs(t time.Time) {
	for _, job := range jobs {
		if job.IsTimeSet(t) {
			host.RunJob(&job)
		}
	}
}
