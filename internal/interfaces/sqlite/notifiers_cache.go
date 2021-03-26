package sqlite

import (
	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// CreateNotifier creates a new job into database
func (c *CacheableWriterReader) CreateNotifier(notifier *entities.Notifier) error {
	return c.wr.CreateNotifier(notifier)
}

// FindNotifiersResponse returns all jobs in FindNotifiersResponse format
func (c *CacheableWriterReader) FindNotifiersResponse() uc.FindNotifiersResponse {
	return c.wr.FindNotifiersResponse()
}
