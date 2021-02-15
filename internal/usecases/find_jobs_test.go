package usecases_test

import (
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

var testsFindJobsResponse = []struct {
	given  string
	expect uc.FindJobsResponse
}{
	{
		given: "every minute",
		expect: uc.FindJobsResponse{
			Jobs: []uc.JobDTO{
				{Name: "list", Command: "ls", Tick: "* * * * * (every minute)"},
			},
			Count: 1,
		},
	},
	{
		given: "* * * * *",
		expect: uc.FindJobsResponse{
			Jobs: []uc.JobDTO{
				{Name: "list", Command: "ls", Tick: "* * * * *"},
			},
			Count: 1,
		},
	},
}

func TestFindJobs(t *testing.T) {
	for _, tt := range testsFindJobsResponse {
		writeReader := mocks.NewStubJobResponseWithExpression(tt.given)
		host := mocks.NewSpyHost()
		uc.New(writeReader, writeReader, host)
		got := uc.FindJobs()
		assertFindJobsResponse(t, got, tt.expect)
	}
}

func assertFindJobsResponse(t *testing.T, got, want uc.FindJobsResponse) {
	if got.Count != want.Count {
		t.Fatalf("got count %v, want count %v", got.Count, want.Count)
	}
	for i := 0; i < len(want.Jobs); i++ {
		if want.Jobs[i] != got.Jobs[i] {
			t.Fatalf("got %v, want %v", got.Jobs[i], want.Jobs[i])
		}
	}
}
