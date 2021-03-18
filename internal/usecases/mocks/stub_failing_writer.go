package mocks

import (
	"errors"

	"github.com/taciomcosta/kronos/internal/entities"
)

// NewFailingWriter stubs a writer that always returns error
func NewFailingWriter() *StubFailingWriter {
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
