package usecases_test

import (
	"errors"
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
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
		reader:   mocks.StubSuccessReader(),
	},
	{
		request:  "non-existing",
		response: uc.DeleteJobResponse{},
		err:      errors.New("resource not found"),
		reader:   mocks.StubFailingReader(),
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
