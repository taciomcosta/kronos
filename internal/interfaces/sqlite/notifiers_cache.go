package sqlite

import (
	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// CreateNotifier creates a new job into database
func (c *CacheableWriterReader) CreateNotifier(notifier *entities.Notifier) error {
	return c.wr.CreateNotifier(notifier)
}

// DeleteNotifier deletes a notifier
func (c *CacheableWriterReader) DeleteNotifier(name string) error {
	return c.wr.DeleteNotifier(name)
}

// FindNotifiersResponse returns all notifiers in FindNotifiersResponse format
func (c *CacheableWriterReader) FindNotifiersResponse() uc.FindNotifiersResponse {
	return c.wr.FindNotifiersResponse()
}

// FindOneNotifier finds all notifiers.
func (c *CacheableWriterReader) FindOneNotifier(name string) (entities.Notifier, error) {
	return c.wr.FindOneNotifier(name)
}
