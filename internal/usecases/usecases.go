package usecases

import (
	"time"

	"github.com/taciomcosta/kronos/internal/entities"
)

var writer Writer
var host entities.Host
var jobs []entities.Job

// New is used for dependency injection on set up.
func New(r Writer, h Host) {
	writer = r
	host = h
}

// GetHost returns host being used
func GetHost() entities.Host {
	return host
}

// Host represents a host where jobs can be run
type Host interface {
	RunJob(job entities.Job)
	GetDettachedStream() entities.Stream
	TickEverySecond() <-chan time.Time
}

// Writer represents a Layer Supertype similar to Writer pattern
// https://martinfowler.com/eaaCatalog/layerSupertype.html
type Writer interface {
	CreateJob(job *entities.Job) error
	FindJobs() []entities.Job
	CountJobs() int
}
