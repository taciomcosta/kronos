package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/taciomcosta/kronos/internal/entities"
)

const slackURL = "https://slack.com/api/chat.postMessage"

type slackPayload struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
}

type slackService struct{}

// Send sends message using chat.postMessage from Slack's API
// https://api.slack.com/methods/chat.postMessage
func (s slackService) Send(msg string, notifier entities.Notifier) error {
	request := newSlackRequest(msg, notifier)
	_, err := http.DefaultClient.Do(request)
	if err != nil {
		return errors.New("error requesting slack notifier " + notifier.Name)
	}
	return nil
}

func newSlackRequest(msg string, notifier entities.Notifier) *http.Request {
	body := newSlackRequestBody(msg, notifier)
	request, _ := http.NewRequest("POST", slackURL, body)
	request.Header.Set("Authorization", "Bearer "+notifier.Metadata["auth_token"])
	request.Header.Set("Content-Type", "application/json")
	return request
}

func newSlackRequestBody(msg string, notifier entities.Notifier) io.Reader {
	payload := slackPayload{Text: msg, Channel: notifier.Metadata["channel_ids"]}
	b, _ := json.Marshal(payload)
	return bytes.NewBuffer(b)
}
