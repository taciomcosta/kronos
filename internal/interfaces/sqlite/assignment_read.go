package sqlite

import (
	"github.com/bvinc/go-sqlite-lite/sqlite3"
	"github.com/taciomcosta/kronos/internal/entities"
)

// FindAssignmentsByJob finds assignments for a job
func (wr *WriterReader) FindAssignmentsByJob(jobName string) []entities.Assignment {
	stmt, _ := db.Prepare(findAssignmentsByJobSQL)
	_ = stmt.Exec(jobName)
	rows := wr.readAllAssignmentRows(stmt)
	return wr.mapAssignmentRowsToEntitities(rows)
}

func (wr *WriterReader) readAllAssignmentRows(stmt *sqlite3.Stmt) []assignmentRow {
	var rows []assignmentRow
	for hasRow, _ := stmt.Step(); hasRow; hasRow, _ = stmt.Step() {
		row := wr.readAssignementRow(stmt)
		rows = append(rows, row)
	}
	return rows
}

func (wr *WriterReader) readAssignementRow(stmt *sqlite3.Stmt) assignmentRow {
	var row assignmentRow
	_ = stmt.Scan(&row.jobName, &row.notifierName, &row.onErrorOnly)
	return row
}

func (wr *WriterReader) mapAssignmentRowsToEntitities(rows []assignmentRow) []entities.Assignment {
	var assignments []entities.Assignment
	for _, row := range rows {
		assignment := entities.Assignment{
			Job:         row.jobName,
			Notifier:    row.notifierName,
			OnErrorOnly: row.onErrorOnly,
		}
		assignments = append(assignments, assignment)
	}
	return assignments
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
	assignment := entities.Assignment{
		Job:         row.jobName,
		Notifier:    row.notifierName,
		OnErrorOnly: row.onErrorOnly,
	}
	return assignment, err
}
