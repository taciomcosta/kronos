package usecases

import (
	"github.com/taciomcosta/kronos/internal/entities"
)

// FindJobs returns a list of all jobs.
func FindJobs() []entities.Job {
	return writer.FindJobs()
}

// CountJobs counts the total of jobs.
func CountJobs() int {
	return writer.CountJobs()
}
