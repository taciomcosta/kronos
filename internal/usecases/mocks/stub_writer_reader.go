package mocks

import (
	"errors"

	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// NewStubWriterReader stubs writer implementation
func NewStubWriterReader() *StubWR {
	return NewStubWriterReaderNJobs(0)
}

// NewStubWriterReaderNJobs stubs writer implementation by creating N jobs
func NewStubWriterReaderNJobs(numJobs int) *StubWR {
	writer := &StubWR{}
	for i := 0; i < numJobs; i++ {
		writer.CreateJobWithExpression("* * * * *")
	}
	return writer
}

// NewStubJobResponseWithExpression stubs FindJobsResponse
func NewStubJobResponseWithExpression(expr string) *StubWR {
	writer := &StubWR{}
	writer.jobsResponse = uc.FindJobsResponse{
		Jobs: []uc.JobDTO{
			{Name: "list", Command: "ls", Tick: expr},
		},
		Count: 1,
	}
	return writer
}

// StubWR implements entities.Writer for tests purposes
type StubWR struct {
	jobs         []entities.Job
	jobsResponse uc.FindJobsResponse
}

// CreateJob creates a job.
func (mr *StubWR) CreateJob(job *entities.Job) error {
	return nil
}

// FindJobs finds all jobs.
func (mr *StubWR) FindJobs() []entities.Job {
	return mr.jobs
}

// FindJobsResponse finds all jobs in FindJobsResponse format
func (mr *StubWR) FindJobsResponse() uc.FindJobsResponse {
	return mr.jobsResponse
}

// CreateJobWithExpression is a shortcut to add a job with provided expression
func (mr *StubWR) CreateJobWithExpression(expression string) {
	job, err := entities.NewJob("name", "cmd", expression, true)
	if err != nil {
		panic(err)
	}
	mr.jobs = append(mr.jobs, job)
}

// CreateDisabledJob is a shortcut to add a job that is disabled
func (mr *StubWR) CreateDisabledJob(expression string) {
	job, err := entities.NewJob("name", "cmd", expression, false)
	if err != nil {
		panic(err)
	}
	mr.jobs = append(mr.jobs, job)
}

// DeleteJob deletes a job
func (mr *StubWR) DeleteJob(name string) error {
	mr.jobs = []entities.Job{}
	return nil
}

// FindOneJob finds one job by name
func (mr *StubWR) FindOneJob(name string) (entities.Job, error) {
	for _, j := range mr.jobs {
		if j.Name == name {
			return j, nil
		}
	}
	return entities.Job{}, errors.New("resource not found")
}

// FindExecutionsResponse finds executions in FindExecution response format
func (mr *StubWR) FindExecutionsResponse(_ uc.FindExecutionsRequest) uc.FindExecutionsResponse {
	return uc.FindExecutionsResponse{
		Executions: []uc.ExecutionDTO{
			{JobName: "list"},
		},
	}
}

// CreateExecution stubs a new Execution creation
func (mr *StubWR) CreateExecution(execution *entities.Execution) error {
	return nil
}

// DescribeJobResponse finds executions in FindExecution response format
func (mr *StubWR) DescribeJobResponse(name string) (uc.DescribeJobResponse, error) {
	return uc.DescribeJobResponse{
		Name:                "list",
		Command:             "ls",
		Tick:                "* * * * *",
		LastExecution:       "2020-01-01T00:00:00.000Z",
		Status:              true,
		ExecutionsSucceeded: 2,
		ExecutionsFailed:    1,
		AverageCPU:          50,
		AverageMem:          1024,
	}, nil
}

// UpdateJob updates a job
func (mr *StubWR) UpdateJob(job *entities.Job) {}

// CreateNotifier creates a notifier
func (mr *StubWR) CreateNotifier(notifier *entities.Notifier) error {
	return nil
}

// FindNotifiersResponse finds all jobs in FindNotifiersResponse format
func (mr *StubWR) FindNotifiersResponse() uc.FindNotifiersResponse {
	return uc.FindNotifiersResponse{
		Count:     1,
		Notifiers: []uc.NotifierDTO{{Name: "myslack", Type: "slack"}},
	}
}
