package mocks

import "github.com/taciomcosta/kronos/internal/entities"

// SpyNotifierService creates a new SpyHost
func SpyNotifierService() *spyNotifierService {
	return &spyNotifierService{}
}

// spyNotifierService is a test double used to spy on notifier services
type spyNotifierService struct {
	called bool
}

// Send sends message to an external notifier service
func (s *spyNotifierService) Send(msg string, notifier entities.Notifier) error {
	s.called = true
	return nil
}

// SendWasCalled checks if Send() was called
func (s *spyNotifierService) SendWasCalled() bool {
	return s.called
}
