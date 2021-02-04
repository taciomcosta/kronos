package domain

var repository Repository
var runner JobsRunner

type Repository interface {
	CreateJob(job *Job) error
	FindJobs() []Job
	CountJobs() int
}

func New(r Repository) {
	repository = r
	go runner.Start()
}
