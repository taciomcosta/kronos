package data

import (
	"fmt"

	"github.com/taciomcosta/kronos/internal/domain"
)

func NewRepository() domain.Repository {
	return &sqliteRepository{}
}

type sqliteRepository struct{}

func (r *sqliteRepository) CreateJob(job *domain.Job) error {
	fmt.Printf("Persisting job %s\n", job.Name)
	return nil
}
