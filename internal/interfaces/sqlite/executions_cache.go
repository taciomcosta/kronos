package sqlite

import (
	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// CreateExecution creates a new job into database
func (c *CacheableWriterReader) CreateExecution(execution *entities.Execution) error {
	return c.wr.CreateExecution(execution)
}

// FindExecutionsResponse returns all jobs in FindExecutionsResponse format
func (c *CacheableWriterReader) FindExecutionsResponse(r uc.FindExecutionsRequest) uc.FindExecutionsResponse {
	return c.wr.FindExecutionsResponse(r)
}
