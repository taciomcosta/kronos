package usecases

import (
	"time"

	"github.com/taciomcosta/kronos/internal/entities"
)

var writer Writer
var reader Reader
var host Host
var notifierService NotifierService

// New is used for dependency injection on set up.
func New(w Writer, r Reader, h Host, n NotifierService) {
	writer = w
	reader = r
	host = h
	notifierService = n
}

// Host represents a host where jobs can be run
type Host interface {
	RunJob(job entities.Job) entities.Execution
	TickEverySecond() <-chan time.Time
}

// Writer represents a Layer Supertype similar to Repository pattern
type Writer interface {
	CreateJob(job *entities.Job) error
	DeleteJob(name string) error
	CreateExecution(execution *entities.Execution) error
	UpdateJob(job *entities.Job)
	CreateNotifier(notifier *entities.Notifier) error
	DeleteNotifier(name string) error
	CreateAssignment(assignment *entities.Assignment) error
}

// Reader represents a Layer Supertype similar to Repository pattern
type Reader interface {
	FindJobs() []entities.Job
	FindOneJob(name string) (entities.Job, error)
	FindJobsResponse() FindJobsResponse
	FindExecutionsResponse(request FindExecutionsRequest) FindExecutionsResponse
	DescribeJobResponse(name string) (DescribeJobResponse, error)
	FindOneNotifier(name string) (entities.Notifier, error)
	FindNotifiersResponse() FindNotifiersResponse
	DescribeNotifierResponse(name string) (DescribeNotifierResponse, error)
}

// NotifierService represents an external service: email, slack, discord,
type NotifierService interface {
	Send(msg string, notifier entities.Notifier) error
}
