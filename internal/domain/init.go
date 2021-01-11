package domain

var repository Repository
var runner JobsRunner

type Repository interface {
	CreateJob(job *Job) error
}

func Init(r Repository) {
	repository = r
	go runner.Start()
}
