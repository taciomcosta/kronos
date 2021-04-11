package mocks

import (
	"time"

	"github.com/taciomcosta/kronos/internal/entities"
)

// StubFailingHost creates a new StubFailingHost
func StubFailingHost() *stubFailingHost {
	stub := stubFailingHost{}
	stub.channel = make(chan time.Time, 1)
	return &stub
}

type stubFailingHost struct {
	channel chan time.Time
}

// RunJob runs a job on stub host
func (s *stubFailingHost) RunJob(job entities.Job) entities.Execution {
	return entities.Execution{
		JobName:  "failing-job",
		Date:     "date",
		Status:   entities.FailedStatus,
		CPUTime:  1000,
		MemUsage: 1000,
	}
}

// TickEverySecond stubs channel so that we can emit desired time on tests
func (s *stubFailingHost) TickEverySecond() <-chan time.Time {
	// In production, we want tick channel to be open forever
	// but we don't this bevahior when testing.
	// Thus we set an expiration time.
	s.expireChannelAfter(10 * time.Millisecond)
	return s.channel
}

func (s *stubFailingHost) expireChannelAfter(duration time.Duration) {
	go func() {
		<-time.After(duration)
		close(s.channel)
	}()
}

// NotifyCurrentTimeIs trigger channel returned by TickEverySecond
func (s *stubFailingHost) NotifyCurrentTimeIs(now time.Time) {
	s.channel <- now
}
