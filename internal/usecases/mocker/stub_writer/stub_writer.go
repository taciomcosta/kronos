package stubwriter

import "github.com/taciomcosta/kronos/internal/entities"

func newStubWriter() *StubWriter {
	stub := &StubWriter{}
	stub.outputs = newDefaultOutputs()
	return stub
}

// StubWriter ...
type StubWriter struct {
	outputs map[string]interface{}
}

// CreateJob creates a job.
func (s *StubWriter) CreateJob(job *entities.Job) error {
	args, _ := s.outputs["CreateJob"].([]interface{})
	return castError(args[0])
}

// DeleteJob deletes a job
func (s *StubWriter) DeleteJob(name string) error {
	args, _ := s.outputs["DeleteJob"].([]interface{})
	return castError(args[0])
}

// CreateExecution stubs a new Execution creation
func (s *StubWriter) CreateExecution(execution *entities.Execution) error {
	args, _ := s.outputs["CreateExecution"].([]interface{})
	return castError(args[0])
}

// UpdateJob updates a job
func (s *StubWriter) UpdateJob(job *entities.Job) {}

// CreateNotifier creates a notifier
func (s *StubWriter) CreateNotifier(notifier *entities.Notifier) error {
	args, _ := s.outputs["CreateNotifier"].([]interface{})
	return castError(args[0])
}

// DeleteNotifier deletes a notifier
func (s *StubWriter) DeleteNotifier(name string) error {
	args, _ := s.outputs["DeleteNotifier"].([]interface{})
	return castError(args[0])
}

// CreateAssignment creates a assignment
func (s *StubWriter) CreateAssignment(assignment *entities.Assignment) error {
	args, _ := s.outputs["CreateAssignment"].([]interface{})
	return castError(args[0])
}

func castError(v interface{}) error {
	if v != nil {
		return v.(error)
	}
	return nil
}
