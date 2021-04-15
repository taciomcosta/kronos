package mocker

import (
	"errors"

	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

func newStubReader(returnFn *ReturnFn) *StubReader {
	returnFn.outputs["FindJobs"] = []interface{}{mustNewJob()}
	stub := &StubReader{returnFn}
	return stub
}

func mustNewJob() entities.Job {
	job, _ := entities.NewJob("name", "cmd", "* * * * *", true)
	return job
}

// StubReader ...
type StubReader struct {
	r *ReturnFn
}

// FindJobs ...
func (s *StubReader) FindJobs() []entities.Job {
	args := s.r.outputs["FindJobs"]
	var output []entities.Job
	for _, a := range args.([]interface{}) {
		output = append(output, a.(entities.Job))
	}
	return output
}

// FindOneJob ...
func (s *StubReader) FindOneJob(name string) (entities.Job, error) {
	args, ok := s.r.outputs["FindOneJob"].([]interface{})
	if !ok {
		return entities.Job{}, errors.New("error")
	}
	return args[0].(entities.Job), args[1].(error)
}

// FindJobsResponse ...
func (s *StubReader) FindJobsResponse() uc.FindJobsResponse {
	arg := s.r.outputs["FindJobsResponse"]
	return arg.(uc.FindJobsResponse)
}

// FindExecutionsResponse ...
func (s *StubReader) FindExecutionsResponse(request uc.FindExecutionsRequest) uc.FindExecutionsResponse {
	arg := s.r.outputs["FindExecutionsResponse"]
	return arg.(uc.FindExecutionsResponse)
}

// DescribeJobResponse ...
func (s *StubReader) DescribeJobResponse(name string) (uc.DescribeJobResponse, error) {
	args, ok := s.r.outputs["FindOneJob"].([]interface{})
	if !ok {
		return uc.DescribeJobResponse{}, errors.New("error")
	}
	return args[0].(uc.DescribeJobResponse), args[1].(error)
}

// FindOneNotifier ...
func (s *StubReader) FindOneNotifier(name string) (entities.Notifier, error) {
	args, ok := s.r.outputs["FindOneNotifier"].([]interface{})
	if !ok {
		return entities.Notifier{}, errors.New("error")
	}
	return args[0].(entities.Notifier), args[1].(error)
}

// FindNotifiersResponse ...
func (s *StubReader) FindNotifiersResponse() uc.FindNotifiersResponse {
	args, ok := s.r.outputs["FindNotifierResponse"].([]interface{})
	if !ok {
		return uc.FindNotifiersResponse{}
	}
	return args[0].(uc.FindNotifiersResponse)
}

// DescribeNotifierResponse ...
func (s *StubReader) DescribeNotifierResponse(name string) (uc.DescribeNotifierResponse, error) {
	args, ok := s.r.outputs["DescribeNotifierResponse"].([]interface{})
	if !ok {
		return uc.DescribeNotifierResponse{}, errors.New("error")
	}
	return args[0].(uc.DescribeNotifierResponse), args[1].(error)
}

// FindAssignmentsByJob ...
func (s *StubReader) FindAssignmentsByJob(jobName string) []entities.Assignment {
	args := s.r.outputs["FindAssignmentsByJob"]
	var output []entities.Assignment
	for _, a := range args.([]interface{}) {
		output = append(output, a.(entities.Assignment))
	}
	return output
}
