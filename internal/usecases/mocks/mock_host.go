package mocks

import (
	"time"

	"github.com/taciomcosta/kronos/internal/entities"
)

// SpyHost is a test double used to spy on host calls
type SpyHost struct {
	called bool
}

// RunJob runs a job on spy host
func (s *SpyHost) RunJob(job *entities.Job) {
	s.called = true
}

// WasCalled tells if RunJob was called
func (s *SpyHost) WasCalled() bool {
	return s.called
}

// GetDettachedStream stubs dettached stream
func (s *SpyHost) GetDettachedStream() entities.Stream {
	return entities.Stream{}
}

// TickEverySecond stubs channel so that we can emit desired time on tests
func (s *SpyHost) TickEverySecond() <-chan time.Time {
	// TODO: stub current time
	ticker := time.NewTicker(1 * time.Second)
	return ticker.C
}
