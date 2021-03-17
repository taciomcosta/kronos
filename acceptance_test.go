package main

import (
	"os"
	"testing"

	"github.com/cucumber/godog"

	"github.com/taciomcosta/kronos/internal/config"
	"github.com/taciomcosta/kronos/internal/interfaces/sqlite"
	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
	"github.com/taciomcosta/kronos/test"
)

var host *mocks.SpyHost

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
		config.EnableTestMode()
		writerReader := sqlite.NewWriterReader(config.GetString("db"))
		host = mocks.NewSpyHost()
		uc.New(writerReader, writerReader, host)
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	jf := features.JobsFeature{}
	ctx.Step(`^I provide valid data for job creation$`, jf.IProvideValidDataForJobCreation)
	ctx.Step(`^I create a new job$`, jf.ICreateANewJob)
	ctx.Step(`^I list the existing jobs$`, jf.IListTheExistingJobs)
	ctx.Step(`^the new job is listed$`, jf.TheNewJobIsListed)
	ctx.Step(`^an error message is shown$`, jf.AnErrorMessageIsShown)
	ctx.Step(`^I provide invalid data for job creation$`, jf.IProvideInvalidDataForJobCreation)
	ctx.Step(`^I delete the new job$`, jf.IDeleteTheNewJob)
	ctx.Step(`^the new job is not listed$`, jf.TheNewJobIsNotListed)
	ctx.Step(`^I describe the new job$`, jf.IDescribeTheNewJob)
	ctx.Step(`^the new job is detailed$`, jf.TheNewJobIsDetailed)

	ef := features.ExecutionsFeature{Host: host}
	ctx.Step(`^(\d+) execution is listed$`, ef.ExecutionIsListed)
	ctx.Step(`^I list all job execution history$`, ef.IListAllJobExecutionHistory)
	ctx.Step(`^that I create a job$`, ef.ThatICreateAJob)
	ctx.Step(`^the job finishes (\d+) execution$`, ef.TheJobFinishesExecution)

}
