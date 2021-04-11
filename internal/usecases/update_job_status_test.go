package usecases_test

import (
	"errors"
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
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
		reader:   mocks.StubSuccessReader(),
	},
	{
		request:  uc.UpdateJobStatusRequest{Name: "name", Status: false},
		response: uc.UpdateJobStatusResponse{Msg: "name disabled"},
		err:      nil,
		reader:   mocks.StubSuccessReader(),
	},
	{
		request:  uc.UpdateJobStatusRequest{Name: "name", Status: false},
		response: uc.UpdateJobStatusResponse{},
		err:      errors.New("resource not found"),
		reader:   mocks.StubFailingReader(),
	},
}

func TestUpdateJobStatus(t *testing.T) {
	for _, tt := range testsUpdateJobStatus {
		writerReader := mocks.StubSuccessWriter()
		host := mocks.NewSpyHost()
		notifierService := mocks.SpyNotifierService()
		uc.New(writerReader, tt.reader, host, notifierService)
		got, err := uc.UpdateJobStatus(tt.request)
		assertEqual(t, got, tt.response)
		assertError(t, err, tt.err)
	}
}
