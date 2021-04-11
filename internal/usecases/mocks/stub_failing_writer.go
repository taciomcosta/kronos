package mocks

import (
	"errors"

	"github.com/taciomcosta/kronos/internal/entities"
)

// StubFailingWriter stubs a writer that always returns error
func StubFailingWriter() *stubFailingWriter {
	return &stubFailingWriter{}
}

type stubFailingWriter struct{}

// CreateJob creates a job.
func (s *stubFailingWriter) CreateJob(job *entities.Job) error {
	return errors.New("StubFailingWriter")
}

// CreateExecution creates a job.
func (s *stubFailingWriter) CreateExecution(e *entities.Execution) error {
	return errors.New("StubFailingWriter")
}

// DeleteJob deletes a job
func (s *stubFailingWriter) DeleteJob(name string) error {
	return errors.New("StubFailingWriter")
}

// UpdateJob updates a job
func (s *stubFailingWriter) UpdateJob(job *entities.Job) {}

// CreateNotifier creates a notifier.
func (s *stubFailingWriter) CreateNotifier(notifier *entities.Notifier) error {
	return errors.New("StubFailingWriter")
}

// DeleteNotifier deletes a notifier
func (s *stubFailingWriter) DeleteNotifier(name string) error {
	return errors.New("StubFailingWriter")
}

// CreateAssignment creates a assignment
func (s *stubFailingWriter) CreateAssignment(assignment *entities.Assignment) error {
	return errors.New("StubFailingWriter")
}
