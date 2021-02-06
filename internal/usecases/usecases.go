package usecases

import "github.com/taciomcosta/kronos/internal/entities"

var repository entities.Repository
var runner entities.JobsRunner

// New is used for dependency injection on set up.
func New(r entities.Repository) {
	repository = r
	startJobsRunner()
}

func startJobsRunner() {
	jobs := repository.FindJobs()
	runner = entities.NewJobRunner(jobs)
	go runner.Start()
}
