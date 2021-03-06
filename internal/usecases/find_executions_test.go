package usecases_test

import (
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocker"
)

func TestFindExecutions(t *testing.T) {
	dependencies := uc.Dependencies{
		mocker.Dependencies().Writer().Build(),
		mocker.Dependencies().Reader().Build(),
		mocker.Dependencies().Host().Build(),
		mocker.Dependencies().NotifierService().Build(),
	}
	uc.New(dependencies)
	got := uc.FindExecutions(uc.FindExecutionsRequest{})
	want := uc.FindExecutionsResponse{
		Executions: []uc.ExecutionDTO{
			{JobName: "list"},
		},
	}
	assertFindExecutionsResponse(t, got, want)
}

func assertFindExecutionsResponse(t *testing.T, got, want uc.FindExecutionsResponse) {
	for i := 0; i < len(want.Executions); i++ {
		if want.Executions[i] != got.Executions[i] {
			t.Fatalf("got %v, want %v", got.Executions[i], want.Executions[i])
		}
	}
}
