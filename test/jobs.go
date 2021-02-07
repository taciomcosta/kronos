package features

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/julienschmidt/httprouter"
	"github.com/taciomcosta/kronos/internal/entities"
	"github.com/taciomcosta/kronos/internal/interfaces/rest"
	"github.com/taciomcosta/kronos/internal/usecases"
)

type JobsFeature struct {
	response *httptest.ResponseRecorder
	inputJob usecases.CreateJobRequest
}

func (j *JobsFeature) IProvideValidDataForJobCreation() error {
	j.inputJob = usecases.CreateJobRequest{
		Name:    "list",
		Command: "ls",
		Tick:    "* * * * *",
	}
	return nil
}

func (j *JobsFeature) IProvideInvalidDataForJobCreation() error {
	j.inputJob = usecases.CreateJobRequest{
		Name:    "list",
		Command: "ls",
		Tick:    "n * * * *",
	}
	return nil
}

func (j *JobsFeature) ICreateANewJob() error {
	request, err := newRequest(j.inputJob)
	j.response = httptest.NewRecorder()
	ps := httprouter.Params{}
	rest.CreateJob(j.response, request, ps)
	return err
}

func newRequest(v interface{}) (*http.Request, error) {
	payload := new(bytes.Buffer)
	err := json.NewEncoder(payload).Encode(v)
	if err != nil {
		return nil, err
	}
	return http.NewRequest("POST", "", payload)
}

func (j *JobsFeature) IListTheExistingJobs() error {
	request, err := http.NewRequest("POST", "", nil)
	j.response = httptest.NewRecorder()
	ps := httprouter.Params{}
	rest.FindJobs(j.response, request, ps)
	return err
}

func (j *JobsFeature) AnErrorMessageIsShown() error {
	var errorMsg rest.ErrorMessage
	err := rest.ReadJSON(j.response.Body, &errorMsg)
	if errorMsg.Msg == "" {
		return errors.New("no error message")
	}
	return err
}

func (j *JobsFeature) TheNewJobShouldBeListed() error {
	var jobs []entities.Job
	err := rest.ReadJSON(j.response.Body, &jobs)
	if err != nil {
		return err
	}
	_, err = findJobByName(jobs, "list")
	return err
}

func findJobByName(jobs []entities.Job, name string) (*entities.Job, error) {
	for _, j := range jobs {
		if j.Name == "list" {
			return &j, nil
		}
	}
	return nil, errors.New("job not listed")
}
