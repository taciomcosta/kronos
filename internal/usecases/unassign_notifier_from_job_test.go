package usecases_test

import (
	"errors"
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocker"
)

var testsUnassignNotifierFromJob = []struct {
	request  uc.UnassignNotifierFromJobRequest
	response uc.UnassignNotifierFromJobResponse
	err      error
	writer   uc.Writer
	reader   uc.Reader
}{
	{
		request: uc.UnassignNotifierFromJobRequest{
			JobName:      "name",
			NotifierName: "myslack",
		},
		response: uc.UnassignNotifierFromJobResponse{
			Msg: "myslack unassigned from name",
		},
		err:    nil,
		writer: mocker.Dependencies().Writer().Build(),
		reader: mocker.Dependencies().Reader().Build(),
	},
	{
		request: uc.UnassignNotifierFromJobRequest{
			JobName:      "name",
			NotifierName: "myslack",
		},
		response: uc.UnassignNotifierFromJobResponse{},
		err:      errors.New("error"),
		writer: mocker.
			Dependencies().Writer().
			Set("DeleteAssignment").
			Return(errors.New("error")).
			Build(),
		reader: mocker.Dependencies().Reader().Build(),
	},
	{
		request: uc.UnassignNotifierFromJobRequest{
			JobName:      "name",
			NotifierName: "myslack",
		},
		response: uc.UnassignNotifierFromJobResponse{},
		err:      errors.New("error"),
		writer:   mocker.Dependencies().Writer().Build(),
		reader: mocker.
			Dependencies().Reader().
			Set("FindOneAssignment").
			Return(mocker.Data().Assignment().Build(), errors.New("error")).
			Build(),
	},
}

func TestUnassignNotifierFromJob(t *testing.T) {
	for _, tt := range testsUnassignNotifierFromJob {
		host := mocker.Dependencies().Host().Build()
		notifierService := mocker.Dependencies().NotifierService().Build()
		dependencies := uc.Dependencies{tt.writer, tt.reader, host, notifierService}
		uc.New(dependencies)
		got, err := uc.UnassignNotifierFromJob(tt.request)
		assertEqual(t, got, tt.response)
		assertError(t, err, tt.err)
	}
}
