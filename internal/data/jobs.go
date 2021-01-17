package data

import (
	"github.com/taciomcosta/kronos/internal/domain"
)

func NewRepository() domain.Repository {
	return &sqliteRepository{}
}

type sqliteRepository struct{}

func (r *sqliteRepository) CreateJob(job *domain.Job) error {
	stmt, err := db.Prepare("INSERT INTO job VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.Exec(nil, job.Name, job.Command, job.Tick)
	return err
}
