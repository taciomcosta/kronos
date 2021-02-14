package main

import (
	"os"
	"testing"

	"github.com/cucumber/godog"

	"github.com/taciomcosta/kronos/internal/interfaces/sqlite"
	"github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
	"github.com/taciomcosta/kronos/test"
)

func TestMain(m *testing.M) {
	status := godog.TestSuite{
		Name:                 "jobs",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
	}.Run()
	os.Exit(status)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		writer := sqlite.NewWriterReader(":memory:")
		host := &mocks.SpyHost{} // TODO: use real host (?)
		usecases.New(writer, writer, host)
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	jf := features.JobsFeature{}
	ctx.Step(`^I provide valid data for job creation$`, jf.IProvideValidDataForJobCreation)
	ctx.Step(`^I create a new job$`, jf.ICreateANewJob)
	ctx.Step(`^I list the existing jobs$`, jf.IListTheExistingJobs)
	ctx.Step(`^the new job should be listed$`, jf.TheNewJobShouldBeListed)
	ctx.Step(`^an error message is shown$`, jf.AnErrorMessageIsShown)
	ctx.Step(`^I provide invalid data for job creation$`, jf.IProvideInvalidDataForJobCreation)
}
