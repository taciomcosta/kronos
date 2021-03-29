package usecases_test

import (
	"errors"
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

func TestDescribeJob(t *testing.T) {
	writer := mocks.NewStubSuccessWriter()
	reader := mocks.NewStubSuccessReader()
	host := mocks.NewSpyHost()
	uc.New(writer, reader, host)
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
	writer := mocks.NewStubSuccessWriter()
	reader := mocks.NewStubFailingReader()
	host := mocks.NewSpyHost()
	uc.New(writer, reader, host)
	got, gotErr := uc.DescribeJob("list")
	want := uc.DescribeJobResponse{}
	wantErr := errors.New("stub-failing-reader")
	assertEqual(t, got, want)
	assertError(t, gotErr, wantErr)
}
