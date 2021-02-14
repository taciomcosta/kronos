package usecases

import (
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
	err = writer.CreateJob(&job)
	if err != nil {
		return CreateJobResponse{}, err
	}
	jobs = append(jobs, job)
	return CreateJobResponse{Msg: job.Name + " created."}, nil
}
