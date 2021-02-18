package mocks

import (
	"errors"

	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// NewStubWriterReader stubs writer implementation
func NewStubWriterReader() *StubWriter {
	return NewStubWriterReaderNJobs(0)
}

// NewStubWriterReaderNJobs stubs writer implementation by creating N jobs
func NewStubWriterReaderNJobs(numJobs int) *StubWriter {
	writer := &StubWriter{}
	for i := 0; i < numJobs; i++ {
		writer.CreateJobWithExpression("* * * * *")
	}
	return writer
}

// NewStubJobResponseWithExpression stubs FindJobsResponse
func NewStubJobResponseWithExpression(expr string) *StubWriter {
	writer := &StubWriter{}
	writer.jobsResponse = uc.FindJobsResponse{
		Jobs: []uc.JobDTO{
			{Name: "list", Command: "ls", Tick: expr},
		},
		Count: 1,
	}
	return writer
}

// StubWriter implements entities.Writer for tests purposes
type StubWriter struct {
	jobs         []entities.Job
	jobsResponse uc.FindJobsResponse
}

// CreateJob creates a job.
func (mr *StubWriter) CreateJob(job *entities.Job) error {
	return nil
}

// FindJobs finds all jobs.
func (mr *StubWriter) FindJobs() []entities.Job {
	return mr.jobs
}

// FindJobsResponse finds all jobs in FindJobsResponse format
func (mr *StubWriter) FindJobsResponse() uc.FindJobsResponse {
	return mr.jobsResponse
}

// CreateJobWithExpression is a shortcut to add a job with provided expression
func (mr *StubWriter) CreateJobWithExpression(expression string) {
	job, err := entities.NewJob("name", "cmd", expression)
	if err != nil {
		panic(err)
	}
	mr.jobs = append(mr.jobs, job)
}

// DeleteJob deletes a job
func (mr *StubWriter) DeleteJob(name string) error {
	mr.jobs = []entities.Job{}
	return nil
}

// FindOneJob finds one job by name
func (mr *StubWriter) FindOneJob(name string) (entities.Job, error) {
	for _, j := range mr.jobs {
		if j.Name == name {
			return j, nil
		}
	}
	return entities.Job{}, errors.New("resource not found")
}
