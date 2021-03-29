package mocks

import (
	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// NewStubSuccessWriter stubs writer implementation
func NewStubSuccessWriter() uc.Writer {
	return &StubW{}
}

// StubW implements entities.Writer for tests purposes
type StubW struct {
	jobs []entities.Job
}

// CreateJob creates a job.
func (mr *StubW) CreateJob(job *entities.Job) error {
	return nil
}

// DeleteJob deletes a job
func (mr *StubW) DeleteJob(name string) error {
	mr.jobs = []entities.Job{}
	return nil
}

// CreateExecution stubs a new Execution creation
func (mr *StubW) CreateExecution(execution *entities.Execution) error {
	return nil
}

// UpdateJob updates a job
func (mr *StubW) UpdateJob(job *entities.Job) {}

// CreateNotifier creates a notifier
func (mr *StubW) CreateNotifier(notifier *entities.Notifier) error {
	return nil
}

// DeleteNotifier deletes a notifier
func (mr *StubW) DeleteNotifier(name string) error {
	return nil
}
