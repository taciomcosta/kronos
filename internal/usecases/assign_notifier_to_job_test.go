package usecases_test

import (
	"errors"
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocker"
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
		writer: mocker.Dependencies().Writer().Build(),
		reader: mocker.Dependencies().Reader().Build(),
	},
	{
		request: uc.AssignNotifierToJobRequest{
			JobName:      "",
			NotifierName: "",
			OnErrorOnly:  true,
		},
		response: uc.AssignNotifierToJobResponse{},
		err:      errors.New("error"),
		writer: mocker.
			Dependencies().Writer().
			Set("CreateAssignment").
			Return(errors.New("error")).Build(),
		reader: mocker.Dependencies().Reader().Build(),
	},
	{
		request: uc.AssignNotifierToJobRequest{
			JobName:      "",
			NotifierName: "",
			OnErrorOnly:  true,
		},
		response: uc.AssignNotifierToJobResponse{},
		err:      errors.New("error finding job/notifier"),
		writer:   mocker.Dependencies().Writer().Build(),
		reader: mocker.
			Dependencies().Reader().
			Set("FindOneJob").
			Return(mocker.Data().Job().Build(), errors.New("error finding job/notifier")).
			Build(),
	},
}

func TestAssignNotifierToJob(t *testing.T) {
	for _, tt := range testsAssignNotifierToJob {
		host := mocker.Dependencies().Host().Build()
		notifierService := mocks.SpyNotifierService()
		dependencies := uc.Dependencies{tt.writer, tt.reader, host, notifierService}
		uc.New(dependencies)
		got, err := uc.AssignNotifierToJob(tt.request)
		assertEqual(t, got, tt.response)
		assertError(t, err, tt.err)
	}
}
