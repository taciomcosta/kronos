package usecases_test

import (
	"errors"
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

var testsAssignNotifierToJob = []struct {
	request  uc.AssignNotifierToJobRequest
	response uc.AssignNotifierToJobResponse
	err      error
	writer   uc.Writer
	reader   uc.Reader
}{
	{
		request: uc.AssignNotifierToJobRequest{
			JobName:      "name",
			NotifierName: "myslack",
			OnErrorOnly:  true,
		},
		response: uc.AssignNotifierToJobResponse{
			Msg: "myslack assigned to name",
		},
		err:    nil,
		writer: mocks.StubSuccessWriter(),
		reader: mocks.NewStubSuccessReader(),
	},
	{
		request: uc.AssignNotifierToJobRequest{
			JobName:      "",
			NotifierName: "",
			OnErrorOnly:  true,
		},
		response: uc.AssignNotifierToJobResponse{},
		err:      errors.New("StubFailingWriter"),
		writer:   mocks.NewStubFailingWriter(),
		reader:   mocks.NewStubSuccessReader(),
	},
	{
		request: uc.AssignNotifierToJobRequest{
			JobName:      "",
			NotifierName: "",
			OnErrorOnly:  true,
		},
		response: uc.AssignNotifierToJobResponse{},
		err:      errors.New("error finding job/notifier"),
		writer:   mocks.StubSuccessWriter(),
		reader:   mocks.NewStubFailingReader(),
	},
}

func TestAssignNotifierToJob(t *testing.T) {
	for _, tt := range testsAssignNotifierToJob {
		host := mocks.NewSpyHost()
		notifierService := mocks.NewSpyNotifierService()
		uc.New(tt.writer, tt.reader, host, notifierService)
		got, err := uc.AssignNotifierToJob(tt.request)
		assertEqual(t, got, tt.response)
		assertError(t, err, tt.err)
	}
}
