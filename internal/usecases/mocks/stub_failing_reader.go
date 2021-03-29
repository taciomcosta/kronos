package mocks

import (
	"errors"

	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// NewFailingReader stubs reader that always fails
func NewStubFailingReader() *StubFailingReader {
	return &StubFailingReader{}
}

// StubFailingReader implements entities.Writer for tests purposes
type StubFailingReader struct{}

// FindJobs finds all jobs.
func (s *StubFailingReader) FindJobs() []entities.Job {
	return []entities.Job{}
}

// FindJobsResponse finds all jobs in FindJobsResponse format
func (s *StubFailingReader) FindJobsResponse() uc.FindJobsResponse {
	return uc.FindJobsResponse{}
}

// FindOneJob finds one job by name
func (s *StubFailingReader) FindOneJob(name string) (entities.Job, error) {
	return entities.Job{}, errors.New("resource not found")
}

// FindExecutionsResponse finds executions in FindExecution response format
func (s *StubFailingReader) FindExecutionsResponse(_ uc.FindExecutionsRequest) uc.FindExecutionsResponse {
	return uc.FindExecutionsResponse{}
}

// DescribeJobResponse finds executions in FindExecution response format
func (s *StubFailingReader) DescribeJobResponse(name string) (uc.DescribeJobResponse, error) {
	return uc.DescribeJobResponse{}, errors.New("stub-failing-reader")
}

// FindNotifiersResponse finds all notifiers in FindNotifiersResponse format
func (s *StubFailingReader) FindNotifiersResponse() uc.FindNotifiersResponse {
	return uc.FindNotifiersResponse{}
}

// FindOneNotifier finds one notifier by name
func (s *StubFailingReader) FindOneNotifier(name string) (entities.Notifier, error) {
	return entities.Notifier{}, errors.New("resource not found")
}
