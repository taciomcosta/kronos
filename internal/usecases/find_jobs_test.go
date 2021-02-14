package usecases_test

import (
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

func TestFindJobs(t *testing.T) {
	writeReader := mocks.NewStubWriterReader()
	host := mocks.NewSpyHost()
	uc.New(writeReader, writeReader, host)
	want := uc.FindJobsResponse{
		Jobs: []uc.JobDTO{
			{Name: "list", Command: "ls", Tick: "* * * * *"},
		},
		Count: 1,
	}
	got := uc.FindJobs()
	assertFindJobsResponse(t, got, want)
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
