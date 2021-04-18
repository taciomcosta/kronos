package usecases_test

import (
	"errors"
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocker"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
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
		reader:   mocker.Stub().Reader().Build(),
	},
	{
		request:  "non-existing",
		response: uc.DeleteJobResponse{},
		err:      errors.New("resource not found"),
		reader: mocker.Stub().Reader().
			Set("FindOneJob").Return(mocker.Data().Job().Build(), errors.New("resource not found")).
			Build(),
	},
}

func TestDeleteJob(t *testing.T) {
	for _, tt := range testsDeleteJob {
		dependencies := uc.Dependencies{
			mocks.StubSuccessWriter(),
			tt.reader,
			mocks.NewSpyHost(),
			mocks.SpyNotifierService(),
		}
		uc.New(dependencies)
		got, err := uc.DeleteJob(tt.request)
		assertEqual(t, got, tt.response)
		assertError(t, err, tt.err)
	}
}
