package entities

// Notifier represents a notifier
type Notifier struct {
	cType    NotifierType
	metadata map[string]string
}

// NotifierType represents a notifier type: slack, email, etc
type NotifierType int

const (
	// SlackChannelType represents slack notifier type
	SlackChannelType NotifierType = 1
)
