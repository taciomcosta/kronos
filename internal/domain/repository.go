package domain

// Repository represents a Layer Supertype for Repository Pattern.
// https://martinfowler.com/eaaCatalog/layerSupertype.html
type Repository interface {
	CreateJob(job *Job) error
	FindJobs() []Job
	CountJobs() int
}
