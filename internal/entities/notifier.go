package entities

// Notifier represents a notifier
type Notifier struct {
	Name     string
	Type     NotifierType
	Metadata map[string]string
}

// NotifierType represents a notifier type: slack, email, etc
type NotifierType int

const (
	// SlackNotifierType represents slack notifier type
	SlackNotifierType NotifierType = 1
)
