package usecases

import (
	"github.com/taciomcosta/kronos/internal/entities"
)

// FindJobsResponse represents response of FindJob usecase
type FindJobsResponse struct {
	Jobs  []JobDTO `json:"jobs"`
	Count int      `json:"count"`
}

// JobDTO represents a Job returned by FindJobsResponse
type JobDTO struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Tick    string `json:"tick"`
	Status  bool   `json:"status"`
}

// FindJobs returns a list of all jobs.
func FindJobs() FindJobsResponse {
	response := reader.FindJobsResponse()
	for i := range response.Jobs {
		response.Jobs[i].Tick = entities.FormatExpression(response.Jobs[i].Tick)
	}
	return response
}
