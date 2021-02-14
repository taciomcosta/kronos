package usecases

import ()

// FindJobsResponse represents response of FindJob usecase
type FindJobsResponse struct {
	Jobs  []JobDTO
	Count int `json:"count"`
}

type JobDTO struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Tick    string `json:"tick"`
}

// FindJobs returns a list of all jobs.
func FindJobs() FindJobsResponse {
	return reader.FindJobsResponse()
}
