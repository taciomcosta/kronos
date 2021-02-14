package mocks

import "github.com/taciomcosta/kronos/internal/entities"

// NewStubWriterReader stubs writer implementation
func NewStubWriterReader() *StubWriter {
	return &StubWriter{}
}

// StubWriter implements entities.Writer for tests purposes
type StubWriter struct {
	jobs []entities.Job
}

// CreateJob creates a job.
func (mr *StubWriter) CreateJob(job *entities.Job) error {
	return nil
}

// CountJobs counts the total of jobs.
func (mr *StubWriter) CountJobs() int {
	return 1
}

// FindJobs finds all jobs.
func (mr *StubWriter) FindJobs() []entities.Job {
	job, _ := entities.NewJob("list", "ls", "* * * * *", entities.Stream{})
	mr.jobs = []entities.Job{job}
	return mr.jobs
}

// CreateJobWithExpression is a shortcut to add a job with provided expression
func (mr *StubWriter) CreateJobWithExpression(expression string) {
	job, _ := entities.NewJob("name", "cmd", expression, entities.Stream{})
	mr.jobs = append(mr.jobs, job)
}
