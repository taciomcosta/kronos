package services

import (
	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// NewNotifierService instantiates a notifier service that sends messages
// using any notifier type
func NewNotifierService() uc.NotifierService {
	return &notifierService{}
}

type notifierService struct {
	slack slackService
}

// Send sends a message using any notifier type
func (n *notifierService) Send(msg string, notifier entities.Notifier) error {
	return n.slack.Send(msg, notifier)
}
