package usecases

import (
	"github.com/taciomcosta/kronos/internal/entities"
)

var repository entities.Repository
var host entities.Host
var jobs []entities.Job

// New is used for dependency injection on set up.
func New(r entities.Repository, h entities.Host) {
	repository = r
	host = h
	ScheduleExistingJobs()
}

// GetHost returns host being used
func GetHost() entities.Host {
	return host
}
