package mocks

import (
	"errors"

	"github.com/taciomcosta/kronos/internal/entities"
)

// NewFailingWriter stubs a writer that always returns error
func NewStubFailingWriter() *StubFailingWriter {
	return &StubFailingWriter{}
}

// StubFailingWriter implements entities.Writer for tests purposes
type StubFailingWriter struct{}

// CreateJob creates a job.
func (s *StubFailingWriter) CreateJob(job *entities.Job) error {
	return errors.New("StubFailingWriter")
}

// CreateExecution creates a job.
func (s *StubFailingWriter) CreateExecution(e *entities.Execution) error {
	return errors.New("StubFailingWriter")
}

// DeleteJob deletes a job
func (s *StubFailingWriter) DeleteJob(name string) error {
	return errors.New("StubFailingWriter")
}

// UpdateJob updates a job
func (s *StubFailingWriter) UpdateJob(job *entities.Job) {}

// CreateNotifier creates a notifier.
func (s *StubFailingWriter) CreateNotifier(notifier *entities.Notifier) error {
	return errors.New("StubFailingWriter")
}

// DeleteNotifier deletes a notifier
func (s *StubFailingWriter) DeleteNotifier(name string) error {
	return errors.New("StubFailingWriter")
}

// CreateAssignment creates a assignment
func (s *StubFailingWriter) CreateAssignment(assignment *entities.Assignment) error {
	return errors.New("StubFailingWriter")
}
