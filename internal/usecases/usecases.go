package usecases

import "github.com/taciomcosta/kronos/internal/domain"

var repository domain.Repository
var runner domain.JobsRunner

// New is used for dependency injection on set up.
func New(r domain.Repository) {
	repository = r
	go runner.Start()
}
