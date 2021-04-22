package usecases_test

import (
	"errors"
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocker"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

func TestDescribeJob(t *testing.T) {
	dependencies := uc.Dependencies{
		mocker.Stub().Writer().Build(),
		mocker.Stub().Reader().Build(),
		mocks.NewSpyHost(),
		mocks.SpyNotifierService(),
	}
	uc.New(dependencies)
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
	dependencies := uc.Dependencies{
		mocker.Stub().Writer().Build(),
		mocker.
			Stub().Reader().
			Set("DescribeJobResponse").
			Return(uc.DescribeJobResponse{}, errors.New("stub-failing-reader")).
			Build(),
		mocks.NewSpyHost(),
		mocks.SpyNotifierService(),
	}
	uc.New(dependencies)
	got, gotErr := uc.DescribeJob("list")
	want := uc.DescribeJobResponse{}
	wantErr := errors.New("stub-failing-reader")
	assertEqual(t, got, want)
	assertError(t, gotErr, wantErr)
}
