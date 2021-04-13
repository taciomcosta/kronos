package mocker

import (
	"errors"

	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// StubReader ...
type StubReader struct {
	r ReturnFn
}

// FindJobs ...
func (s StubReader) FindJobs() []entities.Job {
	jobs, ok := s.r.outputs["FindJobs"]
	if !ok {
		job, _ := entities.NewJob("name", "cmd", "* * * * *", true)
		return []entities.Job{job}
	}
	return jobs.([]entities.Job)
}

// FindOneJob ...
func (s StubReader) FindOneJob(name string) (entities.Job, error) {
	args, ok := s.r.outputs["FindOneJob"].([]interface{})
	if !ok {
		return entities.Job{}, errors.New("error")
	}
	return args[0].(entities.Job), args[1].(error)
}

// FindJobsResponse ...
func (s StubReader) FindJobsResponse() uc.FindJobsResponse {
	arg := s.r.outputs["FindJobsResponse"]
	return arg.(uc.FindJobsResponse)
}

// FindExecutionsResponse ...
func (s StubReader) FindExecutionsResponse(request uc.FindExecutionsRequest) uc.FindExecutionsResponse {
	arg := s.r.outputs["FindExecutionsResponse"]
	return arg.(uc.FindExecutionsResponse)
}

// DescribeJobResponse ...
func (s StubReader) DescribeJobResponse(name string) (uc.DescribeJobResponse, error) {
	args, ok := s.r.outputs["FindOneJob"].([]interface{})
	if !ok {
		return uc.DescribeJobResponse{}, errors.New("error")
	}
	return args[0].(uc.DescribeJobResponse), args[1].(error)
}

// FindOneNotifier ...
func (s StubReader) FindOneNotifier(name string) (entities.Notifier, error) {
	args, ok := s.r.outputs["FindOneNotifier"].([]interface{})
	if !ok {
		return entities.Notifier{}, errors.New("error")
	}
	return args[0].(entities.Notifier), args[1].(error)
}

// FindNotifiersResponse ...
func (s StubReader) FindNotifiersResponse() uc.FindNotifiersResponse {
	args, ok := s.r.outputs["FindNotifierResponse"].([]interface{})
	if !ok {
		return uc.FindNotifiersResponse{}
	}
	return args[0].(uc.FindNotifiersResponse)
}

// DescribeNotifierResponse ...
func (s StubReader) DescribeNotifierResponse(name string) (uc.DescribeNotifierResponse, error) {
	args, ok := s.r.outputs["DescribeNotifierResponse"].([]interface{})
	if !ok {
		return uc.DescribeNotifierResponse{}, errors.New("error")
	}
	return args[0].(uc.DescribeNotifierResponse), args[1].(error)
}

// FindAssignmentsByJob ...
func (s StubReader) FindAssignmentsByJob(jobName string) []entities.Assignment {
	args := s.r.outputs["FindAssignmentsByJob"]
	return args.([]entities.Assignment)
}
