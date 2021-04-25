package spynotifierservice

import (
	"time"

	"github.com/taciomcosta/kronos/internal/entities"
)

func newSpyNotifierService() *SpyNotifierService {
	spy := &SpyNotifierService{}
	spy.outputs = newDefaultOutputs()
	spy.channel = make(chan time.Time, 1)
	return spy
}

// SpyNotifierService is a test double used to spy on notifier services
type SpyNotifierService struct {
	outputs map[string]interface{}
	called  bool
	channel chan time.Time
}

// Send sends message to an external notifier service
func (s *SpyNotifierService) Send(msg string, notifier entities.Notifier) error {
	s.called = true
	args := s.outputs["Send"].([]interface{})
	return castError(args[0])
}

// SendWasCalled checks if Send() was called
func (s *SpyNotifierService) SendWasCalled() bool {
	return s.called
}

func castError(v interface{}) error {
	if v != nil {
		return v.(error)
	}
	return nil
}
