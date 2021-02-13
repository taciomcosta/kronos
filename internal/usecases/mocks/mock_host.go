package mocks

import (
	"time"

	"github.com/taciomcosta/kronos/internal/entities"
)

// NewSpyHost creates a new SpyHost
func NewSpyHost() *SpyHost {
	spy := SpyHost{}
	spy.channel = make(chan time.Time, 1)
	return &spy
}

// SpyHost is a test double used to spy on host calls
type SpyHost struct {
	called  bool
	channel chan time.Time
}

// RunJob runs a job on spy host
func (s *SpyHost) RunJob(job *entities.Job) {
	s.called = true
	close(s.channel)
}

// WasRunJobCalled tells if RunJob was called
func (s *SpyHost) WasRunJobCalled() bool {
	return s.called
}

// GetDettachedStream stubs dettached stream
func (s *SpyHost) GetDettachedStream() entities.Stream {
	return entities.Stream{}
}

// TickEverySecond stubs channel so that we can emit desired time on tests
func (s *SpyHost) TickEverySecond() <-chan time.Time {
	return s.channel
}

// NotifyCurrentTimeIs trigger channel returned by TickEverySecond
func (s *SpyHost) NotifyCurrentTimeIs(now time.Time) {
	select {
	case s.channel <- now:
	case <-time.After(time.Second):
		close(s.channel)
		return
	}
}
