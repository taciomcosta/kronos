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
func (s *SpyHost) RunJob(job entities.Job) {
	s.called = true
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
	// In production, we want tick channel to be open forever
	// but we don't this bevahior when testing.
	// Thus we set an expiration time.
	s.expireChannelAfter(1000 * time.Nanosecond)
	return s.channel
}

func (s *SpyHost) expireChannelAfter(duration time.Duration) {
	go func() {
		<-time.After(duration)
		close(s.channel)
	}()
}

// NotifyCurrentTimeIs trigger channel returned by TickEverySecond
func (s *SpyHost) NotifyCurrentTimeIs(now time.Time) {
	s.channel <- now
}
