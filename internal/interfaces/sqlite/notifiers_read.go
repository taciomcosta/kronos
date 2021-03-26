package sqlite

import (
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// FindNotifiersResponse returns all jobs in FindNotifiersResponse format
func (wr *WriterReader) FindNotifiersResponse() uc.FindNotifiersResponse {
	return uc.FindNotifiersResponse{}
}
