package mocker

import (
	"errors"

	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

func newStubReader() *StubReader {
	stub := &StubReader{}
	stub.outputs = newDefaultOutputs()
	return stub
}

func newDefaultOutputs() map[string]interface{} {
	var outputs = make(map[string]interface{})
	d := &defaultStubReader{}
	outputs["FindJobs"] = d.FindJobs()
	outputs["FindOneJob"] = d.FindOneJob()
	outputs["FindJobsResponse"] = d.FindJobsResponse()
	outputs["FindExecutionsResponse"] = d.FindExecutionsResponse()
	outputs["DescribeJobResponse"] = d.DescribeJobResponse()
	outputs["FindOneNotifier"] = d.FindOneNotifier()
	outputs["FindNotifiersResponse"] = d.FindNotifiersResponse()
	outputs["DescribeNotifierResponse"] = d.DescribeNotifierResponse()
	outputs["FindAssignmentsByJob"] = d.FindAssignmentsByJob()
	return outputs
}

// StubReader ...
type StubReader struct {
	outputs map[string]interface{}
}

// FindJobs ...
func (s *StubReader) FindJobs() []entities.Job {
	args := s.outputs["FindJobs"]
	var output []entities.Job
	for _, a := range args.([]interface{}) {
		output = append(output, a.(entities.Job))
	}
	return output
}

// FindOneJob ...
func (s *StubReader) FindOneJob(name string) (entities.Job, error) {
	args, ok := s.outputs["FindOneJob"].([]interface{})
	if !ok {
		return entities.Job{}, errors.New("error")
	}
	return args[0].(entities.Job), args[1].(error)
}

// FindJobsResponse ...
func (s *StubReader) FindJobsResponse() uc.FindJobsResponse {
	arg := s.outputs["FindJobsResponse"]
	return arg.(uc.FindJobsResponse)
}

// FindExecutionsResponse ...
func (s *StubReader) FindExecutionsResponse(request uc.FindExecutionsRequest) uc.FindExecutionsResponse {
	args := s.outputs["FindExecutionsResponse"].([]interface{})
	return args[0].(uc.FindExecutionsResponse)
}

// DescribeJobResponse ...
func (s *StubReader) DescribeJobResponse(name string) (uc.DescribeJobResponse, error) {
	args, ok := s.outputs["FindOneJob"].([]interface{})
	if !ok {
		return uc.DescribeJobResponse{}, errors.New("error")
	}
	return args[0].(uc.DescribeJobResponse), args[1].(error)
}

// FindOneNotifier ...
func (s *StubReader) FindOneNotifier(name string) (entities.Notifier, error) {
	args, _ := s.outputs["FindOneNotifier"].([]interface{})
	var err error = nil
	if args[1] != nil {
		err = args[1].(error)
	}
	return args[0].(entities.Notifier), err
}

// FindNotifiersResponse ...
func (s *StubReader) FindNotifiersResponse() uc.FindNotifiersResponse {
	args, ok := s.outputs["FindNotifierResponse"].([]interface{})
	if !ok {
		return uc.FindNotifiersResponse{}
	}
	return args[0].(uc.FindNotifiersResponse)
}

// DescribeNotifierResponse ...
func (s *StubReader) DescribeNotifierResponse(name string) (uc.DescribeNotifierResponse, error) {
	args, ok := s.outputs["DescribeNotifierResponse"].([]interface{})
	if !ok {
		return uc.DescribeNotifierResponse{}, errors.New("error")
	}
	return args[0].(uc.DescribeNotifierResponse), args[1].(error)
}

// FindAssignmentsByJob ...
func (s *StubReader) FindAssignmentsByJob(jobName string) []entities.Assignment {
	args := s.outputs["FindAssignmentsByJob"]
	var output []entities.Assignment
	for _, a := range args.([]interface{}) {
		output = append(output, a.(entities.Assignment))
	}
	return output
}
