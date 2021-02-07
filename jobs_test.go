package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/julienschmidt/httprouter"

	"github.com/taciomcosta/kronos/internal/entities"
	"github.com/taciomcosta/kronos/internal/interfaces/rest"
	"github.com/taciomcosta/kronos/internal/interfaces/sqlite"
	"github.com/taciomcosta/kronos/internal/usecases"
)

func TestMain(m *testing.M) {
	status := godog.TestSuite{
		Name:                 "jobs",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
	}.Run()
	os.Exit(status)
}

type jobsFeature struct {
	response *httptest.ResponseRecorder
}

func (j *jobsFeature) iCreateANewJob() error {
	useCaseRequest := usecases.CreateJobRequest{
		Name:    "list",
		Command: "ls",
		Tick:    "* * * * *",
	}
	payload := new(bytes.Buffer)
	err := json.NewEncoder(payload).Encode(useCaseRequest)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", "", payload)
	if err != nil {
		return err
	}

	response := httptest.NewRecorder()

	ps := httprouter.Params{}

	rest.CreateJob(response, request, ps)
	return err
}

func (j *jobsFeature) iListTheExistingJobs() error {
	request, err := http.NewRequest("POST", "", nil)

	j.response = httptest.NewRecorder()

	ps := httprouter.Params{}

	rest.FindJobs(j.response, request, ps)
	return err
}

func (j *jobsFeature) theNewJobShouldBeListed() error {
	decoder := json.NewDecoder(j.response.Body)
	var jobs []entities.Job
	err := decoder.Decode(&jobs)
	if err != nil {
		return err
	}
	for _, j := range jobs {
		if j.Name == "list" {
			return nil
		}
	}
	return errors.New("job not listed")
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		repository := sqlite.NewRepository(":memory:")
		usecases.New(repository)
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	jf := &jobsFeature{}
	ctx.Step(`^I create a new job$`, jf.iCreateANewJob)
	ctx.Step(`^I list the existing jobs$`, jf.iListTheExistingJobs)
	ctx.Step(`^the new job should be listed$`, jf.theNewJobShouldBeListed)
}
