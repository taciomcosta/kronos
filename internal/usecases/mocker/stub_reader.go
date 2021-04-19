package mocker

import (
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
	args, _ := s.outputs["FindOneJob"].([]interface{})
	return args[0].(entities.Job), castError(args[1])
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
	args, _ := s.outputs["FindOneJob"].([]interface{})
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

func castError(v interface{}) error {
	if v != nil {
		return v.(error)
	}
	return nil
}
