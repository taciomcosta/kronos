package mocks

import (
	"errors"

	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// NewFailingWriter stubs a writer that always returns error
func NewFailingWriter() *StubFailingWriter {
	return &StubFailingWriter{}
}

// StubFailingWriter implements entities.Writer for tests purposes
type StubFailingWriter struct {
	jobs         []entities.Job
	jobsResponse uc.FindJobsResponse
}

// CreateJob creates a job.
func (mr *StubFailingWriter) CreateJob(job *entities.Job) error {
	return errors.New("StubFailingWriter")
}

// DeleteJob deletes a job
func (mr *StubFailingWriter) DeleteJob(name string) error {
	return errors.New("StubFailingWriter")
}
