package mocks

import "github.com/taciomcosta/kronos/internal/entities"

// NewMockRepository returns a mock implementation of repository.
func NewMockRepository() *MockRepository {
	return &MockRepository{}
}

// MockRepository implements entities.Repository for tests purposes
type MockRepository struct {
	jobs []entities.Job
}

// CreateJob creates a job.
func (mr *MockRepository) CreateJob(job *entities.Job) error {
	return nil
}

// CountJobs counts the total of jobs.
func (mr *MockRepository) CountJobs() int {
	return 1
}

// FindJobs finds all jobs.
func (mr *MockRepository) FindJobs() []entities.Job {
	job, _ := entities.NewJob("list", "ls", "* * * * *", entities.Stream{})
	mr.jobs = []entities.Job{job}
	return mr.jobs
}

// CreateJobWithExpression is a shortcut to add a job with provided expression
func (mr *MockRepository) CreateJobWithExpression(expression string) {
	job, _ := entities.NewJob("name", "cmd", expression, entities.Stream{})
	mr.jobs = append(mr.jobs, job)
}
