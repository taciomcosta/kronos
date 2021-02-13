package mocks

import "github.com/taciomcosta/kronos/internal/entities"

// NewStubRepository subts repository implementation
func NewStubRepository() *StubRepository {
	return &StubRepository{}
}

// StubRepository implements entities.Repository for tests purposes
type StubRepository struct {
	jobs []entities.Job
}

// CreateJob creates a job.
func (mr *StubRepository) CreateJob(job *entities.Job) error {
	return nil
}

// CountJobs counts the total of jobs.
func (mr *StubRepository) CountJobs() int {
	return 1
}

// FindJobs finds all jobs.
func (mr *StubRepository) FindJobs() []entities.Job {
	job, _ := entities.NewJob("list", "ls", "* * * * *", entities.Stream{})
	mr.jobs = []entities.Job{job}
	return mr.jobs
}

// CreateJobWithExpression is a shortcut to add a job with provided expression
func (mr *StubRepository) CreateJobWithExpression(expression string) {
	job, _ := entities.NewJob("name", "cmd", expression, entities.Stream{})
	mr.jobs = append(mr.jobs, job)
}
