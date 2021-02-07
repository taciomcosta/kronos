package main

import (
	"github.com/cucumber/godog"
)

type jobsFeature struct {
}

func (j *jobsFeature) iCreateANewJob() error {
	return nil
}

func (j *jobsFeature) iListTheExistingJobs() error {
	return nil
}

func (j *jobsFeature) theNewJobShouldBeListed() error {
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	jobs := &jobsFeature{}
	ctx.Step(`^I create a new job$`, jobs.iCreateANewJob)
	ctx.Step(`^I list the existing jobs$`, jobs.iListTheExistingJobs)
	ctx.Step(`^the new job should be listed$`, jobs.theNewJobShouldBeListed)
}
