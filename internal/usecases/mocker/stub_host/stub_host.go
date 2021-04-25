package stubhost

import (
	"time"

	"github.com/taciomcosta/kronos/internal/entities"
)

func newStubHost() *StubHost {
	stub := &StubHost{}
	stub.outputs = newDefaultOutputs()
	return stub
}

// StubHost ...
type StubHost struct {
	outputs map[string]interface{}
	channel chan time.Time
}

// RunJob ...
func (s *StubHost) RunJob(job entities.Job) entities.Execution {
	args := s.outputs["RunJob"].([]interface{})
	return args[0].(entities.Execution)
}

// TickEverySecond stubs channel so that we can emit desired time on tests
func (s *StubHost) TickEverySecond() <-chan time.Time {
	// In production, we want tick channel to be open forever
	// but we don't this bevahior when testing.
	// Thus we set an expiration time.
	s.expireChannelAfter(10 * time.Millisecond)
	return s.channel
}

func (s *StubHost) expireChannelAfter(duration time.Duration) {
	go func() {
		<-time.After(duration)
		close(s.channel)
	}()
}

// NotifyCurrentTimeIs trigger channel returned by TickEverySecond
func (s *StubHost) NotifyCurrentTimeIs(now time.Time) {
	s.channel <- now
}
