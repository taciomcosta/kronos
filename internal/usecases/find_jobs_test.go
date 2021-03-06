package usecases_test

import (
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocker"
)

var testsFindJobsResponse = []struct {
	given  string
	expect uc.FindJobsResponse
}{
	{
		given: "every minute",
		expect: uc.FindJobsResponse{
			Jobs: []uc.JobDTO{
				{
					Name:    "name",
					Command: "cmd",
					Tick:    "* * * * * (every minute)",
					Status:  true,
				},
			},
			Count: 1,
		},
	},
	{
		given: "* * * * *",
		expect: uc.FindJobsResponse{
			Jobs: []uc.JobDTO{
				{
					Name:    "name",
					Command: "cmd",
					Tick:    "* * * * *",
					Status:  true,
				},
			},
			Count: 1,
		},
	},
}

func TestFindJobs(t *testing.T) {
	for _, tt := range testsFindJobsResponse {
		dependencies := uc.Dependencies{
			mocker.Dependencies().Writer().Build(),
			mocker.
				Dependencies().Reader().
				Set("FindJobsResponse").
				Return(mocker.Data().FindJobsResponse().WithExpression(tt.given).Build()).
				Build(),
			mocker.Dependencies().Host().Build(),
			mocker.Dependencies().NotifierService().Build(),
		}
		uc.New(dependencies)
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
