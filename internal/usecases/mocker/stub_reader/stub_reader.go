package stubreader

import (
	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

func newStubReader() *StubReader {
	stub := &StubReader{}
	stub.outputs = newDefaultOutputs()
	return stub
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
	args, _ := s.outputs["FindOneJob"].([]interface{})
	return args[0].(entities.Job), castError(args[1])
}

// FindJobsResponse ...
func (s *StubReader) FindJobsResponse() uc.FindJobsResponse {
	args, _ := s.outputs["FindJobsResponse"].([]interface{})
	return args[0].(uc.FindJobsResponse)
}

// FindExecutionsResponse ...
func (s *StubReader) FindExecutionsResponse(request uc.FindExecutionsRequest) uc.FindExecutionsResponse {
	args := s.outputs["FindExecutionsResponse"].([]interface{})
	return args[0].(uc.FindExecutionsResponse)
}

// DescribeJobResponse ...
func (s *StubReader) DescribeJobResponse(name string) (uc.DescribeJobResponse, error) {
	args, _ := s.outputs["DescribeJobResponse"].([]interface{})
	return args[0].(uc.DescribeJobResponse), castError(args[1])

}

// FindOneNotifier ...
func (s *StubReader) FindOneNotifier(name string) (entities.Notifier, error) {
	args, _ := s.outputs["FindOneNotifier"].([]interface{})
	return args[0].(entities.Notifier), castError(args[1])
}

// FindNotifiersResponse ...
func (s *StubReader) FindNotifiersResponse() uc.FindNotifiersResponse {
	args, _ := s.outputs["FindNotifiersResponse"].([]interface{})
	return args[0].(uc.FindNotifiersResponse)
}

// DescribeNotifierResponse ...
func (s *StubReader) DescribeNotifierResponse(name string) (uc.DescribeNotifierResponse, error) {
	args, _ := s.outputs["DescribeNotifierResponse"].([]interface{})
	return args[0].(uc.DescribeNotifierResponse), castError(args[1])
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

// FindOneAssignment ...
func (s *StubReader) FindOneAssignment(jobName string, notifierName string) (entities.Assignment, error) {
	args, _ := s.outputs["FindOneAssignment"].([]interface{})
	return args[0].(entities.Assignment), castError(args[1])
}

func castError(v interface{}) error {
	if v != nil {
		return v.(error)
	}
	return nil
}
