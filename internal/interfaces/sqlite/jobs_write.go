package sqlite

import (
	"github.com/taciomcosta/kronos/internal/entities"
)

// CreateJob creates a job.
func (wr *WriterReader) CreateJob(job *entities.Job) error {
	stmt, err := db.Prepare("INSERT INTO job VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	err = stmt.Exec(job.Name, job.Command, job.Tick)
	_ = stmt.Close()
	return err
}
