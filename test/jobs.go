package features

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/julienschmidt/httprouter"
	"github.com/taciomcosta/kronos/internal/interfaces/rest"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// JobsFeature contains BDD steps related to jobs feature
type JobsFeature struct {
	responseFindJobs    *httptest.ResponseRecorder
	responseDescribeJob *httptest.ResponseRecorder
	inputJob            uc.CreateJobRequest
}

// IProvideValidDataForJobCreation represents a BDD step
func (j *JobsFeature) IProvideValidDataForJobCreation() error {
	j.inputJob = uc.CreateJobRequest{
		Name:    "list",
		Command: "ls",
		Tick:    "* * * * *",
	}
	return nil
}

// IProvideInvalidDataForJobCreation represents a BDD step
func (j *JobsFeature) IProvideInvalidDataForJobCreation() error {
	j.inputJob = uc.CreateJobRequest{
		Name:    "list",
		Command: "ls",
		Tick:    "n * * * *",
	}
	return nil
}

// ICreateANewJob represents a BDD step
func (j *JobsFeature) ICreateANewJob() error {
	request, err := newRequest(j.inputJob)
	j.responseFindJobs = httptest.NewRecorder()
	ps := httprouter.Params{}
	rest.CreateJob(j.responseFindJobs, request, ps)
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

// IListTheExistingJobs represents a BDD step
func (j *JobsFeature) IListTheExistingJobs() error {
	request, err := http.NewRequest("GET", "", nil)
	j.responseFindJobs = httptest.NewRecorder()
	ps := httprouter.Params{}
	rest.FindJobs(j.responseFindJobs, request, ps)
	return err
}

// AnErrorMessageIsShown represents a BDD step
func (j *JobsFeature) AnErrorMessageIsShownForJob() error {
	var errorMsg rest.ErrorMessage
	err := rest.ReadJSON(j.responseFindJobs.Body, &errorMsg)
	if errorMsg.Msg == "" {
		return errors.New("no error message")
	}
	return err
}

// TheNewJobIsListed represents a BDD step
func (j *JobsFeature) TheNewJobIsListed() error {
	var findJobsResponse uc.FindJobsResponse
	err := rest.ReadJSON(j.responseFindJobs.Body, &findJobsResponse)
	if err != nil {
		return err
	}
	job := findJobByName(findJobsResponse, "list")
	if job == nil {
		return errors.New("job not listed when it should")
	}
	return nil
}

func findJobByName(response uc.FindJobsResponse, name string) *uc.JobDTO {
	for _, j := range response.Jobs {
		if j.Name == "list" {
			return &j
		}
	}
	return nil
}

// IDeleteTheNewJob represents a BDD step
func (j *JobsFeature) IDeleteTheNewJob() error {
	request, err := http.NewRequest("DELETE", "", nil)
	response := httptest.NewRecorder()
	name := httprouter.Param{Key: "name", Value: j.inputJob.Name}
	params := httprouter.Params{name}
	rest.DeleteJob(response, request, params)
	return err
}

// TheNewJobIsNotListed represents a BDD step
func (j *JobsFeature) TheNewJobIsNotListed() error {
	var response uc.FindJobsResponse
	err := rest.ReadJSON(j.responseFindJobs.Body, &response)
	if err != nil {
		return err
	}
	job := findJobByName(response, "list")
	if job != nil {
		return errors.New("job was listed when it should not")
	}
	return nil
}

// IDescribeTheNewJob represents a BDD step
func (j *JobsFeature) IDescribeTheNewJob() error {
	request, err := http.NewRequest("GET", "", nil)
	j.responseDescribeJob = httptest.NewRecorder()
	name := httprouter.Param{Key: "name", Value: j.inputJob.Name}
	ps := httprouter.Params{name}
	rest.DescribeJob(j.responseDescribeJob, request, ps)
	return err
}

// TheNewJobIsDetailed represents a BDD step
func (j *JobsFeature) TheNewJobIsDetailed() error {
	var response uc.DescribeJobResponse
	err := rest.ReadJSON(j.responseDescribeJob.Body, &response)
	if err != nil {
		return err
	}
	if response.Name != j.inputJob.Name {
		return errors.New("job not described")
	}
	return nil
}
