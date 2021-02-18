package usecases

import (
	"time"

	"github.com/taciomcosta/kronos/internal/entities"
)

var writer Writer
var reader Reader
var host entities.Host
var jobs []entities.Job

// New is used for dependency injection on set up.
func New(w Writer, r Reader, h Host) {
	writer = w
	reader = r
	host = h
}

// GetHost returns host being used
func GetHost() entities.Host {
	return host
}

// Host represents a host where jobs can be run
type Host interface {
	RunJob(job entities.Job)
	TickEverySecond() <-chan time.Time
}

// Writer represents a Layer Supertype similar to Repository pattern
type Writer interface {
	CreateJob(job *entities.Job) error
	DeleteJob(name string) error
}

// Reader represents a Layer Supertype similar to Repository pattern
type Reader interface {
	FindJobs() []entities.Job
	FindOneJob(name string) (entities.Job, error)
	FindJobsResponse() FindJobsResponse
}
