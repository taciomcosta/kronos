package usecases_test

import (
	"errors"
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

func TestDescribeJob(t *testing.T) {
	writer := mocks.StubSuccessWriter()
	reader := mocks.NewStubSuccessReader()
	host := mocks.NewSpyHost()
	notifierService := mocks.NewSpyNotifierService()
	uc.New(writer, reader, host, notifierService)
	got, err := uc.DescribeJob("list")
	want := uc.DescribeJobResponse{
		Name:                "list",
		Command:             "ls",
		Tick:                "* * * * *",
		LastExecution:       "2020-01-01T00:00:00.000Z",
		Status:              true,
		ExecutionsSucceeded: 2,
		ExecutionsFailed:    1,
		AverageCPU:          50,
		AverageMem:          1024,
	}
	assertEqual(t, got, want)
	assertError(t, err, nil)
}

func TestDescribeJobFailure(t *testing.T) {
	writer := mocks.StubSuccessWriter()
	reader := mocks.NewStubFailingReader()
	host := mocks.NewSpyHost()
	notifierService := mocks.NewSpyNotifierService()
	uc.New(writer, reader, host, notifierService)
	got, gotErr := uc.DescribeJob("list")
	want := uc.DescribeJobResponse{}
	wantErr := errors.New("stub-failing-reader")
	assertEqual(t, got, want)
	assertError(t, gotErr, wantErr)
}
