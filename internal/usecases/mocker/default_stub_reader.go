package mocker

import (
	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// defaultStubReader implements entities.Reader for tests purposes
type defaultStubReader struct{}

// FindJobs finds all jobs.
func (mr *defaultStubReader) FindJobs() []interface{} {
	job, _ := entities.NewJob("name", "cmd", "* * * * *", true)
	return []interface{}{job}
}

// FindJobsResponse finds all jobs in FindJobsResponse format
func (mr *defaultStubReader) FindJobsResponse() uc.FindJobsResponse {
	return uc.FindJobsResponse{
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
}

// FindOneJob finds one job by name
func (mr *defaultStubReader) FindOneJob(name string) (entities.Job, error) {
	return entities.NewJob("name", "cmd", "* * * * *", true)
}

// FindExecutionsResponse finds executions in FindExecution response format
func (mr *defaultStubReader) FindExecutionsResponse(_ uc.FindExecutionsRequest) uc.FindExecutionsResponse {
	return uc.FindExecutionsResponse{
		Executions: []uc.ExecutionDTO{
			{JobName: "list"},
		},
	}
}

// DescribeJobResponse finds executions in FindExecution response format
func (mr *defaultStubReader) DescribeJobResponse(name string) (uc.DescribeJobResponse, error) {
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
func (mr *defaultStubReader) FindNotifiersResponse() uc.FindNotifiersResponse {
	return uc.FindNotifiersResponse{
		Count:     1,
		Notifiers: []uc.NotifierDTO{{Name: "myslack", Type: "slack"}},
	}
}

// FindOneNotifier finds one notifier by name
func (mr *defaultStubReader) FindOneNotifier(name string) (entities.Notifier, error) {
	return entities.Notifier{Name: "myslack"}, nil
}

// DescribeNotifierResponse finds executions in FindExecution response format
func (mr *defaultStubReader) DescribeNotifierResponse(name string) (uc.DescribeNotifierResponse, error) {
	return uc.DescribeNotifierResponse{
		Name: "myslack",
		Type: "slack",
		Metadata: map[string]string{
			"auth_token":  "123",
			"channel_ids": "1,2,3",
		},
	}, nil
}

// FindFindAssignmentsByJob finds assignments for a job
func (mr *defaultStubReader) FindAssignmentsByJob(jobName string) []entities.Assignment {
	return []entities.Assignment{
		{
			Job:         "name",
			Notifier:    "myslack",
			OnErrorOnly: false,
		},
	}
}
