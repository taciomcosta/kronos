package mocks

import "github.com/taciomcosta/kronos/internal/entities"

// SpyHost is a test double used to spy on host calls
type SpyHost struct {
	called bool
}

// RunJob runs a job on spy host
func (s *SpyHost) RunJob(job *entities.Job) error {
	s.called = true
	return nil
}

// WasCalled tells if RunJob was called
func (s *SpyHost) WasCalled() bool {
	return s.called
}

// GetDettachedStream stubs dettached stream
func (s *SpyHost) GetDettachedStream() entities.Stream {
	return entities.Stream{}
}
