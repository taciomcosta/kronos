package entities

import (
	"errors"
	"fmt"
	"strings"
)

// NewNotifier creates a new notifier
func NewNotifier(name string, ntype NotifierType, metadata map[string]string) (Notifier, error) {
	err := validate(name, ntype, metadata)
	if err != nil {
		return Notifier{}, err
	}
	return Notifier{Name: name, Type: ntype, Metadata: metadata}, nil
}

func validate(name string, ntype NotifierType, metadata map[string]string) error {
	_, ok := expectedMetadata[ntype]
	if !ok {
		return errors.New("invalid notifier type")
	}
	missing := findMissingMetadata(metadata, expectedMetadata[ntype])
	if len(missing) > 0 {
		return fmt.Errorf("expected %s to be provided", strings.Join(missing, ", "))
	}
	return nil
}

func findMissingMetadata(given map[string]string, expected []string) []string {
	missing := []string{}
	for _, key := range expected {
		_, ok := given[key]
		if !ok {
			missing = append(missing, key)
		}
	}
	return missing
}

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

var expectedMetadata = map[NotifierType][]string{
	SlackNotifierType: {"auth_token", "channel_ids"},
}
