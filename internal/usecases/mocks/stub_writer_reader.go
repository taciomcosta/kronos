package mocks

import (
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

// StubWriter implements entities.Writer for tests purposes
type StubWriter struct {
	jobs []entities.Job
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
	response := uc.FindJobsResponse{
		Jobs: []uc.JobDTO{
			{Name: "list", Command: "ls", Tick: "* * * * *"},
		},
		Count: 1,
	}
	return response
}

// CreateJobWithExpression is a shortcut to add a job with provided expression
func (mr *StubWriter) CreateJobWithExpression(expression string) {
	job, err := entities.NewJob("name", "cmd", expression, entities.Stream{})
	if err != nil {
		panic(err)
	}
	mr.jobs = append(mr.jobs, job)
}
