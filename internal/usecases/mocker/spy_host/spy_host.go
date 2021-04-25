package spyhost

import (
	"time"

	"github.com/taciomcosta/kronos/internal/entities"
)

func newSpyHost() *SpyHost {
	spy := &SpyHost{}
	spy.outputs = newDefaultOutputs()
	return spy
}

// SpyHost ...
type SpyHost struct {
	outputs map[string]interface{}
	called  bool
	channel chan time.Time
}

// RunJob ...
func (s *SpyHost) RunJob(job entities.Job) entities.Execution {
	s.called = true
	args := s.outputs["RunJob"].([]interface{})
	return args[0].(entities.Execution)
}

// DidJobRun tells if RunJob was called
func (s *SpyHost) DidJobRun() bool {
	return s.called
}

// TickEverySecond spys channel so that we can emit desired time on tests
func (s *SpyHost) TickEverySecond() <-chan time.Time {
	// In production, we want tick channel to be open forever
	// but we don't this bevahior when testing.
	// Thus we set an expiration time.
	s.expireChannelAfter(10 * time.Millisecond)
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
