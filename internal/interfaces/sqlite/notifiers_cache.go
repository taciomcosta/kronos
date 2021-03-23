package sqlite

import (
	"github.com/taciomcosta/kronos/internal/entities"
)

// CreateNotifier creates a new job into database
func (c *CacheableWriterReader) CreateNotifier(notifier *entities.Notifier) error {
	return c.wr.CreateNotifier(notifier)
}
