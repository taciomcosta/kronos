package usecases

import (
	"github.com/taciomcosta/kronos/internal/domain"
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
	job, err := domain.NewJob(request.Name, request.Command, request.Tick)
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
func FindJobs() []domain.Job {
	return repository.FindJobs()
}

// CountJobs counts the total of jobs.
func CountJobs() int {
	return repository.CountJobs()
}
