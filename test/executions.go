package features

import (
	//"bytes"
	//"encoding/json"
	//"errors"
	//"net/http"

	//"github.com/julienschmidt/httprouter"
	//"github.com/taciomcosta/kronos/internal/interfaces/rest"
	"net/http/httptest"

	"github.com/cucumber/godog"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// ExecutionsFeature contains BDD steps related to executions feature
type ExecutionsFeature struct {
	response *httptest.ResponseRecorder
	inputJob uc.CreateJobRequest
}

// ExecutionShouldIsListed contains BDD steps related to executions feature
func (ef *ExecutionsFeature) ExecutionShouldIsListed(arg1 int) error {
	return godog.ErrPending
}

// IListAllJobExecutionHistory contains BDD steps related to executions feature
func (ef *ExecutionsFeature) IListAllJobExecutionHistory() error {
	return godog.ErrPending
}

// ThatICreateAJob contains BDD steps related to executions feature
func (ef *ExecutionsFeature) ThatICreateAJob() error {
	return godog.ErrPending
}

// TheJobFinishesExecution contains BDD steps related to executions feature
func (ef *ExecutionsFeature) TheJobFinishesExecution(arg1 int) error {
	return godog.ErrPending
}
