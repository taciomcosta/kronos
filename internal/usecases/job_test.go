// +build unit !integration

package usecases

import (
	"errors"
	"testing"

	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

var testsCreateJob = []struct {
	request  CreateJobRequest
	response CreateJobResponse
	err      error
}{
	{
		request: CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "* * * * *",
		},
		response: CreateJobResponse{
			Msg: "list created.",
		},
		err: nil,
	},
	{
		request: CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "n * * * *",
		},
		response: CreateJobResponse{},
		err:      errors.New("can't parse n"),
	},
	{
		request: CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "* n * * *",
		},
		response: CreateJobResponse{},
		err:      errors.New("can't parse n"),
	},
	{
		request: CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "* * n * *",
		},
		response: CreateJobResponse{},
		err:      errors.New("can't parse n"),
	},
	{
		request: CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "* * * n *",
		},
		response: CreateJobResponse{},
		err:      errors.New("can't parse n"),
	},
	{
		request: CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "* * * * n",
		},
		response: CreateJobResponse{},
		err:      errors.New("can't parse n"),
	},
}

func TestCreateJob(t *testing.T) {
	New(mocks.NewMockRepository(), &mocks.SpyHost{})
	for _, tt := range testsCreateJob {
		response, err := CreateJob(tt.request)
		if tt.response != response {
			t.Errorf("got %v, expected %v", response, tt.response)
		}
		assertError(t, err, tt.err)
	}
}

func assertError(t *testing.T, got error, want error) {
	if got == nil && want == nil {
		return
	}
	if got == nil && want != nil {
		t.Errorf("expected error %v, got %v", want, got)
	}
	if got != nil && want == nil {
		t.Errorf("expected error %v, got %v", want, got)
	}
	if got.Error() != want.Error() {
		t.Errorf("expected error %v, got %v", want, got)
	}
}
