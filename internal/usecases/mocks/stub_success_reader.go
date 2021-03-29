package mocks

import (
	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// NewStubSuccessReader stubs success reader
func NewStubSuccessReader() uc.Reader {
	return NewStubSuccessReaderWithExpr("* * * * *")
}

// NewStubSuccessReaderWithDisabledJob stubs success reader with disabled job
func NewStubSuccessReaderWithDisabledJob(expression string) uc.Reader {
	job, _ := entities.NewJob("name", "cmd", expression, false)
	return &StubR{job}
}

// NewStubSuccessReaderWithExpr stubs success reader with job expr
func NewStubSuccessReaderWithExpr(expression string) uc.Reader {
	job, _ := entities.NewJob("name", "cmd", expression, true)
	return &StubR{job}
}

// StubR implements entities.Reader for tests purposes
type StubR struct {
	job entities.Job
}

// FindJobs finds all jobs.
func (mr *StubR) FindJobs() []entities.Job {
	return []entities.Job{mr.job}
}

// FindJobsResponse finds all jobs in FindJobsResponse format
func (mr *StubR) FindJobsResponse() uc.FindJobsResponse {
	return uc.FindJobsResponse{
		Count: 1,
		Jobs: []uc.JobDTO{
			{

				Name:    mr.job.Name,
				Command: mr.job.Command,
				Tick:    mr.job.Tick,
				Status:  mr.job.Status,
			},
		},
	}
}

// FindOneJob finds one job by name
func (mr *StubR) FindOneJob(name string) (entities.Job, error) {
	return mr.job, nil
}

// FindExecutionsResponse finds executions in FindExecution response format
func (mr *StubR) FindExecutionsResponse(_ uc.FindExecutionsRequest) uc.FindExecutionsResponse {
	return uc.FindExecutionsResponse{
		Executions: []uc.ExecutionDTO{
			{JobName: "list"},
		},
	}
}

// DescribeJobResponse finds executions in FindExecution response format
func (mr *StubR) DescribeJobResponse(name string) (uc.DescribeJobResponse, error) {
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

// FindNotifiersResponse finds all notifiers in FindNotifiersResponse format
func (mr *StubR) FindNotifiersResponse() uc.FindNotifiersResponse {
	return uc.FindNotifiersResponse{
		Count:     1,
		Notifiers: []uc.NotifierDTO{{Name: "myslack", Type: "slack"}},
	}
}

// FindOneNotifier finds one notifier by name
func (mr *StubR) FindOneNotifier(name string) (entities.Notifier, error) {
	return entities.Notifier{Name: "myslack"}, nil
}
