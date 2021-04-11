package mocks

import (
	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// StubSuccessWriter stubs writer implementation
func StubSuccessWriter() uc.Writer {
	return &stubSuccessWriter{}
}

// stubSuccessWriter implements entities.Writer for tests purposes
type stubSuccessWriter struct {
	jobs []entities.Job
}

// CreateJob creates a job.
func (mr *stubSuccessWriter) CreateJob(job *entities.Job) error {
	return nil
}

// DeleteJob deletes a job
func (mr *stubSuccessWriter) DeleteJob(name string) error {
	mr.jobs = []entities.Job{}
	return nil
}

// CreateExecution stubs a new Execution creation
func (mr *stubSuccessWriter) CreateExecution(execution *entities.Execution) error {
	return nil
}

// UpdateJob updates a job
func (mr *stubSuccessWriter) UpdateJob(job *entities.Job) {}

// CreateNotifier creates a notifier
func (mr *stubSuccessWriter) CreateNotifier(notifier *entities.Notifier) error {
	return nil
}

// DeleteNotifier deletes a notifier
func (mr *stubSuccessWriter) DeleteNotifier(name string) error {
	return nil
}

// CreateAssignment creates a assignment
func (mr *stubSuccessWriter) CreateAssignment(assignment *entities.Assignment) error {
	return nil
}
