package mocks

import (
	"errors"

	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// StubFailingReader stubs reader that always fails
func StubFailingReader() *stubFailingReader {
	return &stubFailingReader{}
}

type stubFailingReader struct{}

// FindJobs finds all jobs.
func (s *stubFailingReader) FindJobs() []entities.Job {
	return []entities.Job{}
}

// FindJobsResponse finds all jobs in FindJobsResponse format
func (s *stubFailingReader) FindJobsResponse() uc.FindJobsResponse {
	return uc.FindJobsResponse{}
}

// FindOneJob finds one job by name
func (s *stubFailingReader) FindOneJob(name string) (entities.Job, error) {
	return entities.Job{}, errors.New("resource not found")
}

// FindExecutionsResponse finds executions in FindExecution response format
func (s *stubFailingReader) FindExecutionsResponse(_ uc.FindExecutionsRequest) uc.FindExecutionsResponse {
	return uc.FindExecutionsResponse{}
}

// DescribeJobResponse finds executions in FindExecution response format
func (s *stubFailingReader) DescribeJobResponse(name string) (uc.DescribeJobResponse, error) {
	return uc.DescribeJobResponse{}, errors.New("stub-failing-reader")
}

// FindNotifiersResponse finds all notifiers in FindNotifiersResponse format
func (s *stubFailingReader) FindNotifiersResponse() uc.FindNotifiersResponse {
	return uc.FindNotifiersResponse{}
}

// FindOneNotifier finds one notifier by name
func (s *stubFailingReader) FindOneNotifier(name string) (entities.Notifier, error) {
	return entities.Notifier{}, errors.New("resource not found")
}

// DescribeNotifierResponse finds executions in FindExecution response format
func (s *stubFailingReader) DescribeNotifierResponse(name string) (uc.DescribeNotifierResponse, error) {
	return uc.DescribeNotifierResponse{}, errors.New("stub-failing-reader")
}
