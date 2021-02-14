package usecases

import (
	"errors"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
	"testing"
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
		response: CreateJobResponse{Msg: "list created."},
		err:      nil,
	},
	{
		request: CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "1 1 1 1 1",
		},
		response: CreateJobResponse{Msg: "list created."},
		err:      nil,
	},
	{
		request: CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "1,2,3 1,2,3 1,2,3 1,2,3 1,2,3",
		},
		response: CreateJobResponse{Msg: "list created."},
		err:      nil,
	},
	{
		request: CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "1,*,3 1,*,3 1,*,3 1,*,3 1,*,3",
		},
		response: CreateJobResponse{Msg: "list created."},
		err:      nil,
	},
	{
		request: CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "1-4 1-4 1-4 1-4 1-4",
		},
		response: CreateJobResponse{Msg: "list created."},
		err:      nil,
	},
	{
		request: CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "*/2 */2 */2 */2 */2",
		},
		response: CreateJobResponse{Msg: "list created."},
		err:      nil,
	},
	{
		request: CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "1-5/2 1-5/2 1-5/2 1-5/2 1-5/2",
		},
		response: CreateJobResponse{Msg: "list created."},
		err:      nil,
	},
	{
		request: CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "1-5/2,6 1-5/2,6 1-5/2,6 1-5/2,6 1-5/2,6",
		},
		response: CreateJobResponse{Msg: "list created."},
		err:      nil,
	},
	{
		request:  CreateJobRequest{Tick: "n * * * *"},
		response: CreateJobResponse{},
		err:      errors.New("can't parse n"),
	},
	{
		request:  CreateJobRequest{Tick: "* n * * *"},
		response: CreateJobResponse{},
		err:      errors.New("can't parse n"),
	},
	{
		request:  CreateJobRequest{Tick: "* * n * *"},
		response: CreateJobResponse{},
		err:      errors.New("can't parse n"),
	},
	{
		request:  CreateJobRequest{Tick: "* * * n *"},
		response: CreateJobResponse{},
		err:      errors.New("can't parse n"),
	},
	{
		request:  CreateJobRequest{Tick: "* * * * n"},
		response: CreateJobResponse{},
		err:      errors.New("can't parse n"),
	},
	{
		request:  CreateJobRequest{Tick: "60 * * * *"},
		response: CreateJobResponse{},
		err:      errors.New("60 out of range 0-59"),
	},
	{
		request:  CreateJobRequest{Tick: "1- * * * *"},
		response: CreateJobResponse{},
		err:      errors.New("missing value in range 1-"),
	},
	{
		request:  CreateJobRequest{Tick: "0-60 * * * *"},
		response: CreateJobResponse{},
		err:      errors.New("0-60 out of range 0-59"),
	},
	{
		request:  CreateJobRequest{Tick: "* * 0-23 * *"},
		response: CreateJobResponse{},
		err:      errors.New("0-23 out of range 1-31"),
	},
	{
		request:  CreateJobRequest{Tick: "* * * 1-13/2,7 *"},
		response: CreateJobResponse{},
		err:      errors.New("1-13 out of range 1-12"),
	},
	{
		request:  CreateJobRequest{Tick: "* * * 1-5/2,13 *"},
		response: CreateJobResponse{},
		err:      errors.New("13 out of range 1-12"),
	},
	{
		request:  CreateJobRequest{Tick: "10/2 * * * *"},
		response: CreateJobResponse{},
		err:      errors.New("can't parse 10/2"),
	},
}

func TestCreateJob(t *testing.T) {
	writerReader := mocks.NewStubWriterReader()
	New(writerReader, writerReader, mocks.NewSpyHost())
	for _, tt := range testsCreateJob {
		response, err := CreateJob(tt.request)
		assertResponse(t, response, tt.response)
		assertError(t, err, tt.err)
	}
}

func assertResponse(t *testing.T, got CreateJobResponse, want CreateJobResponse) {
	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	if got == nil && want == nil {
		return
	}
	if got == nil && want != nil {
		t.Fatalf("expected error %v, got %v", want, got)
	}
	if got != nil && want == nil {
		t.Fatalf("expected error %v, got %v", want, got)
	}
	if got.Error() != want.Error() {
		t.Fatalf("expected error %v, got %v", want, got)
	}
}
