package domain

var repository Repository
var runner JobsRunner

// Repository represents a Layer Supertype for Repository Pattern.
// https://martinfowler.com/eaaCatalog/layerSupertype.html
type Repository interface {
	CreateJob(job *Job) error
	FindJobs() []Job
	CountJobs() int
}

// New is used for dependency injection on set up.
func New(r Repository) {
	repository = r
	go runner.Start()
}
