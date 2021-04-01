package usecases_test

import (
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

func TestFindExecutions(t *testing.T) {
	writer := mocks.NewStubSuccessWriter()
	reader := mocks.NewStubSuccessReader()
	host := mocks.NewSpyHost()
	notifierService := mocks.NewSpyNotifierService()
	uc.New(writer, reader, host, notifierService)
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
