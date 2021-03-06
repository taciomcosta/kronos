package usecases_test

import (
	"errors"
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocker"
)

var testsUpdateJobStatus = []struct {
	request  uc.UpdateJobStatusRequest
	response uc.UpdateJobStatusResponse
	err      error
	reader   uc.Reader
}{
	{
		request:  uc.UpdateJobStatusRequest{Name: "name", Status: true},
		response: uc.UpdateJobStatusResponse{Msg: "name enabled"},
		err:      nil,
		reader:   mocker.Dependencies().Reader().Build(),
	},
	{
		request:  uc.UpdateJobStatusRequest{Name: "name", Status: false},
		response: uc.UpdateJobStatusResponse{Msg: "name disabled"},
		err:      nil,
		reader:   mocker.Dependencies().Reader().Build(),
	},
	{
		request:  uc.UpdateJobStatusRequest{Name: "name", Status: false},
		response: uc.UpdateJobStatusResponse{},
		err:      errors.New("resource not found"),
		reader: mocker.
			Dependencies().Reader().
			Set("FindOneJob").
			Return(mocker.Data().Job().Build(), errors.New("resource not found")).
			Build(),
	},
}

func TestUpdateJobStatus(t *testing.T) {
	for _, tt := range testsUpdateJobStatus {
		dependencies := uc.Dependencies{
			mocker.Dependencies().Writer().Build(),
			tt.reader,
			mocker.Dependencies().Host().Build(),
			mocker.Dependencies().NotifierService().Build(),
		}
		uc.New(dependencies)
		got, err := uc.UpdateJobStatus(tt.request)
		assertEqual(t, got, tt.response)
		assertError(t, err, tt.err)
	}
}
