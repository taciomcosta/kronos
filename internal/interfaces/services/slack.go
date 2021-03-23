package services

import (
	"fmt"

	"github.com/taciomcosta/kronos/internal/entities"
)

type slackService struct{}

func (s slackService) Send(msg string, notifier entities.Notifier) error {
	fmt.Println("sending slack message")
	return nil
}
