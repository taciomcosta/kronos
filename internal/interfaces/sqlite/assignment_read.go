package sqlite

import "github.com/taciomcosta/kronos/internal/entities"

// FindAssignmentsByJob finds assignments for a job
func (wr *WriterReader) FindAssignmentsByJob(jobName string) []entities.Assignment {
	return []entities.Assignment{}
}

// FindOneAssignment finds one assignment
func (wr *WriterReader) FindOneAssignment(jobName string, notifierName string) (entities.Assignment, error) {
	var row assignmentRow
	stmt, _ := db.Prepare(findOneAssignmentSQL)
	_ = stmt.Exec(jobName, notifierName)
	hasRow, _ := stmt.Step()
	if !hasRow {
		return entities.Assignment{}, errResourceNotFound
	}
	err := stmt.Scan(&row.jobName, &row.notifierName, &row.onErrorOnly)
	return entities.Assignment{
		Job:         row.jobName,
		Notifier:    row.notifierName,
		OnErrorOnly: row.onErrorOnly,
	}, err
}
