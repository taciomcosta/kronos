package mocker

import "github.com/taciomcosta/kronos/internal/entities"

func newNotifierBuilder() *NotifierBuilder {
	notifier := entities.Notifier{
		Name: "myslack",
		Type: entities.SlackNotifierType,
		Metadata: map[string]string{
			"auth_token":  "123",
			"channel_ids": "1,2,3",
		},
	}
	builder := &NotifierBuilder{notifier}
	return builder
}

// NotifierBuilder is Tests Data Builder Pattern for entities.Notifier
type NotifierBuilder struct {
	assignment entities.Notifier
}

// Build builds a new entities.Notifier
func (b *NotifierBuilder) Build() entities.Notifier {
	return b.assignment
}
