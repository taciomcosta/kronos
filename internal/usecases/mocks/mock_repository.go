package mocks

import "github.com/taciomcosta/kronos/internal/entities"

// NewMockRepository returns a mock implementation of repository.
func NewMockRepository() entities.Repository {
	return &mockRepository{}
}

type mockRepository struct{}

// CreateJob creates a job.
func (mr *mockRepository) CreateJob(job *entities.Job) error {
	return nil
}

// CountJobs counts the total of jobs.
func (mr *mockRepository) CountJobs() int {
	return 1
}

// FindJobs finds all jobs.
func (mr *mockRepository) FindJobs() []entities.Job {
	return []entities.Job{
		{
			Name:    "list",
			Command: "ls",
			Tick:    "* * * * *",
		},
	}

}
