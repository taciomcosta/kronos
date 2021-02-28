package sqlite

import (
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// FindExecutionsResponse returns all jobs in FindExecutionsResponse format
func (c *CacheableWriterReader) FindExecutionsResponse(r uc.FindExecutionsRequest) uc.FindExecutionsResponse {
	return c.wr.FindExecutionsResponse(r)
}
