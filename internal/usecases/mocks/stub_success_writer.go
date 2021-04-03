package mocks

import (
	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// NewStubSuccessWriter stubs writer implementation
func NewStubSuccessWriter() uc.Writer {
	return &StubSuccessWriter{}
}

// StubSuccessWriter implements entities.Writer for tests purposes
type StubSuccessWriter struct {
	jobs []entities.Job
}

// CreateJob creates a job.
func (mr *StubSuccessWriter) CreateJob(job *entities.Job) error {
	return nil
}

// DeleteJob deletes a job
func (mr *StubSuccessWriter) DeleteJob(name string) error {
	mr.jobs = []entities.Job{}
	return nil
}

// CreateExecution stubs a new Execution creation
func (mr *StubSuccessWriter) CreateExecution(execution *entities.Execution) error {
	return nil
}

// UpdateJob updates a job
func (mr *StubSuccessWriter) UpdateJob(job *entities.Job) {}

// CreateNotifier creates a notifier
func (mr *StubSuccessWriter) CreateNotifier(notifier *entities.Notifier) error {
	return nil
}

// DeleteNotifier deletes a notifier
func (mr *StubSuccessWriter) DeleteNotifier(name string) error {
	return nil
}

// CreateAssignment creates a assignment
func (mr *StubSuccessWriter) CreateAssignment(assignment *entities.Assignment) error {
	return nil
}
