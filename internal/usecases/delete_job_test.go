package usecases_test

import (
	"errors"
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocker"
)

var testsDeleteJob = []struct {
	request  string
	response uc.DeleteJobResponse
	err      error
	reader   uc.Reader
}{
	{
		request:  "name",
		response: uc.DeleteJobResponse{Msg: "name deleted"},
		err:      nil,
		reader:   mocker.Dependencies().Reader().Build(),
	},
	{
		request:  "non-existing",
		response: uc.DeleteJobResponse{},
		err:      errors.New("resource not found"),
		reader: mocker.Dependencies().Reader().
			Set("FindOneJob").Return(mocker.Data().Job().Build(), errors.New("resource not found")).
			Build(),
	},
}

func TestDeleteJob(t *testing.T) {
	for _, tt := range testsDeleteJob {
		dependencies := uc.Dependencies{
			mocker.Dependencies().Writer().Build(),
			tt.reader,
			mocker.Dependencies().Host().Build(),
			mocker.Dependencies().NotifierService().Build(),
		}
		uc.New(dependencies)
		got, err := uc.DeleteJob(tt.request)
		assertEqual(t, got, tt.response)
		assertError(t, err, tt.err)
	}
}
