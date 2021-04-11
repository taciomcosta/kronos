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
		notifierService := mocks.SpyNotifierService()
		uc.New(writerReader, writerReader, host, notifierService)
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	jf := features.JobsFeature{}
	ctx.Step(`^I provide valid data for job creation$`, jf.IProvideValidDataForJobCreation)
	ctx.Step(`^I create a new job$`, jf.ICreateANewJob)
	ctx.Step(`^I list the existing jobs$`, jf.IListTheExistingJobs)
	ctx.Step(`^the new job is listed$`, jf.TheNewJobIsListed)
	ctx.Step(`^an error message is shown for job$`, jf.AnErrorMessageIsShownForJob)
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

	nf := features.NotifiersFeature{}
	ctx.Step(`^I provide valid data for notifier creation$`, nf.IProvideValidDataForNotifierCreation)
	ctx.Step(`^I create a new notifier$`, nf.ICreateANewNotifier)
	ctx.Step(`^I list the existing notifiers$`, nf.IListTheExistingNotifiers)
	ctx.Step(`^the new notifier is listed$`, nf.TheNewNotifierIsListed)
	ctx.Step(`^an error message is shown for notifier$`, nf.AnErrorMessageIsShownForNotifier)
	ctx.Step(`^I provide invalid data for notifier creation$`, nf.IProvideInvalidDataForNotifierCreation)
	ctx.Step(`^I delete the new notifier$`, nf.IDeleteTheNewNotifier)
	ctx.Step(`^the new notifier is not listed$`, nf.TheNewNotifierIsNotListed)
	ctx.Step(`^I describe the new notifier$`, nf.IDescribeTheNewNotifier)
	ctx.Step(`^the new notifier is detailed$`, nf.TheNewNotifierIsDetailed)
}
