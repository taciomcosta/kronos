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
		reader:   mocks.NewStubSuccessReader(),
	},
	{
		request:  "non-existing",
		response: uc.DeleteJobResponse{},
		err:      errors.New("resource not found"),
		reader:   mocks.NewStubFailingReader(),
	},
}

func TestDeleteJob(t *testing.T) {
	for _, tt := range testsDeleteJob {
		writer := mocks.StubSuccessWriter()
		host := mocks.NewSpyHost()
		notifierService := mocks.NewSpyNotifierService()
		uc.New(writer, tt.reader, host, notifierService)
		got, err := uc.DeleteJob(tt.request)
		assertEqual(t, got, tt.response)
		assertError(t, err, tt.err)
	}
}
