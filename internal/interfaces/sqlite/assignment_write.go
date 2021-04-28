package sqlite

import (
	"github.com/taciomcosta/kronos/internal/entities"
)

// CreateAssignment creates a new assignment into database
func (wr *WriterReader) CreateAssignment(assignment *entities.Assignment) error {
	return wr.runWriteOperation(
		insertAssignmentSQL,
		assignment.Job,
		assignment.Notifier,
		assignment.OnErrorOnly,
	)
}

// DeleteAssignment deletes a notifier
func (wr *WriterReader) DeleteAssignment(assignment *entities.Assignment) error {
	return wr.runWriteOperation(
		deleteAssignmentSQL,
		assignment.Job,
		assignment.Notifier,
	)
}
