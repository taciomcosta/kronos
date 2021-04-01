package mocks

import "github.com/taciomcosta/kronos/internal/entities"

// NewSpyNotifierService creates a new SpyHost
func NewSpyNotifierService() *SpyNotifierService {
	return &SpyNotifierService{}
}

// SpyNotifierService is a test double used to spy on notifier services
type SpyNotifierService struct {
	called bool
}

// Send sends message to an external notifier service
func (s *SpyNotifierService) Send(msg string, notifier entities.Notifier) error {
	s.called = true
	return nil
}
