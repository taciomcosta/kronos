package sqlite

import (
	"github.com/taciomcosta/kronos/internal/entities"
)

// CreateJob creates a new job into database
func (wr *WriterReader) CreateJob(job *entities.Job) error {
	return wr.runWriteOperation(insertJobSQL, job.Name, job.Command, job.Tick, job.Status)
}

// DeleteJob deletes a job
func (wr *WriterReader) DeleteJob(name string) error {
	return wr.runWriteOperation(deleteJobSQL, name)
}

// UpdateJob updates a job
func (wr *WriterReader) UpdateJob(job *entities.Job) {
	_ = wr.runWriteOperation(updateJobSQL, job.Status, job.Name)
}

func (wr *WriterReader) runWriteOperation(sql string, args ...interface{}) error {
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	err = stmt.Exec(args...)
	_ = stmt.Close()
	if err != nil {
		return parseError(err)
	}
	return nil
}
