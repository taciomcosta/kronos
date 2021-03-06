package sqlite

import "github.com/taciomcosta/kronos/internal/entities"

// CreateExecution creates a new job into database
func (wr *WriterReader) CreateExecution(execution *entities.Execution) error {
	return wr.runWriteOperation(
		insertExecutionSQL,
		execution.JobName,
		execution.Date,
		execution.Status,
		execution.MemUsage,
		execution.CPUTime,
	)
}
