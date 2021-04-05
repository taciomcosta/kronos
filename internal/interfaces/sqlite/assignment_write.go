package sqlite

import (
	"github.com/taciomcosta/kronos/internal/entities"
)

// CreateAssignment creates a new assignment into database
func (wr *WriterReader) CreateAssignment(assignment *entities.Assignment) error {
	return wr.runWriteOperation(
		insertAssignmentSQL,
		assignment.Job.Name,
		assignment.Notifier.Name,
		assignment.OnErrorOnly,
	)
}
