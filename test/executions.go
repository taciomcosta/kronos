package features

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/taciomcosta/kronos/internal/interfaces/rest"
	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

// ExecutionsFeature contains BDD steps related to executions feature
type ExecutionsFeature struct {
	Host     *mocks.SpyHost
	response *httptest.ResponseRecorder
}

// ThatICreateAJob contains BDD steps related to executions feature
func (e *ExecutionsFeature) ThatICreateAJob() error {
	jobRequest := uc.CreateJobRequest{
		Name:    "list",
		Command: "ls",
		Tick:    "* * * * *",
	}
	request, err := newRequest(jobRequest)
	e.response = httptest.NewRecorder()
	ps := httprouter.Params{}
	rest.CreateJob(e.response, request, ps)
	return err
}

// TheJobFinishesExecution contains BDD steps related to executions feature
func (e *ExecutionsFeature) TheJobFinishesExecution(arg1 int) error {
	currentTime := time.Date(2021, 1, 13, 0, 0, 0, 0, time.UTC)
	e.Host.NotifyCurrentTimeIs(currentTime)
	uc.ScheduleExistingJobs()
	if !e.Host.DidJobRun() {
		return errors.New("job did not execute")
	}
	return nil
}

// IListAllJobExecutionHistory contains BDD steps related to executions feature
func (e *ExecutionsFeature) IListAllJobExecutionHistory() error {
	request, err := http.NewRequest("GET", "", nil)
	e.response = httptest.NewRecorder()
	ps := httprouter.Params{}
	rest.FindExecutions(e.response, request, ps)
	return err
}

// ExecutionIsListed contains BDD steps related to executions feature
func (e *ExecutionsFeature) ExecutionIsListed(arg1 int) error {
	var findExecutionsResponse uc.FindExecutionsResponse
	err := rest.ReadJSON(e.response.Body, &findExecutionsResponse)
	if err != nil {
		return err
	}
	if len(findExecutionsResponse.Executions) != 1 {
		return errors.New("execution not listed when it should")
	}
	return nil
}
