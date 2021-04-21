package stubreader

import (
	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocker/data"
)

// defaultStubReader implements entities.Reader for tests purposes
type defaultStubReader struct{}

// FindJobs finds all jobs.
func (mr *defaultStubReader) FindJobs() []interface{} {
	job, _ := entities.NewJob("name", "cmd", "* * * * *", true)
	return []interface{}{job}
}

// FindJobsResponse finds all jobs in FindJobsResponse format
func (mr *defaultStubReader) FindJobsResponse() []interface{} {
	response := uc.FindJobsResponse{
		Count: 1,
		Jobs: []uc.JobDTO{
			{

				Name:    "name",
				Command: "cmd",
				Tick:    "* * * * *",
				Status:  true,
			},
		},
	}
	return []interface{}{response}
}

// FindOneJob finds one job by name
func (mr *defaultStubReader) FindOneJob() []interface{} {
	job, err := entities.NewJob("name", "cmd", "* * * * *", true)
	return []interface{}{job, err}
}

// FindExecutionsResponse finds executions in FindExecution response format
func (mr *defaultStubReader) FindExecutionsResponse() []interface{} {
	response := uc.FindExecutionsResponse{
		Executions: []uc.ExecutionDTO{
			{JobName: "list"},
		},
	}
	return []interface{}{response}
}

// DescribeJobResponse finds executions in FindExecution response format
func (mr *defaultStubReader) DescribeJobResponse() []interface{} {
	response := uc.DescribeJobResponse{
		Name:                "list",
		Command:             "ls",
		Tick:                "* * * * *",
		LastExecution:       "2020-01-01T00:00:00.000Z",
		Status:              true,
		ExecutionsSucceeded: 2,
		ExecutionsFailed:    1,
		AverageCPU:          50,
		AverageMem:          1024,
	}
	return []interface{}{response, nil}
}

// FindNotifiersResponse finds all notifiers in FindNotifiersResponse format
func (mr *defaultStubReader) FindNotifiersResponse() []interface{} {
	response := uc.FindNotifiersResponse{
		Count:     1,
		Notifiers: []uc.NotifierDTO{{Name: "myslack", Type: "slack"}},
	}
	return []interface{}{response}
}

// FindOneNotifier finds one notifier by name
func (mr *defaultStubReader) FindOneNotifier() []interface{} {
	notifier := entities.Notifier{Name: "myslack"}
	return []interface{}{notifier, nil}
}

// DescribeNotifierResponse finds executions in FindExecution response format
func (mr *defaultStubReader) DescribeNotifierResponse() []interface{} {
	response := uc.DescribeNotifierResponse{
		Name: "myslack",
		Type: "slack",
		Metadata: map[string]string{
			"auth_token":  "123",
			"channel_ids": "1,2,3",
		},
	}
	return []interface{}{response, nil}
}

// FindFindAssignmentsByJob finds assignments for a job
func (mr *defaultStubReader) FindAssignmentsByJob() []interface{} {
	dataMocker := &data.DataMocker{}
	assignment := dataMocker.Assignment().Build()
	return []interface{}{assignment}
}
