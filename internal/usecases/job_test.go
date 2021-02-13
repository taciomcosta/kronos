package usecases

import (
	"errors"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
	"testing"
	"time"
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
	New(mocks.NewMockRepository(), mocks.NewSpyHost())
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

var testScheduleExistingJobs = []struct {
	expression string
	times      []time.Time
	isSet      bool
}{
	{
		expression: "* * * * *",
		times: []time.Time{
			time.Date(2021, 1, 1, 1, 1, 1, 0, time.UTC),
			time.Date(2021, 1, 1, 2, 1, 1, 0, time.UTC),
		},
		isSet: true,
	},
	{
		expression: "*/2 * * * *",
		times: []time.Time{
			time.Date(2021, 1, 1, 1, 1, 1, 0, time.UTC),
			time.Date(2021, 1, 1, 1, 3, 1, 0, time.UTC),
			time.Date(2021, 1, 1, 1, 7, 1, 0, time.UTC),
		},
		isSet: false,
	},
	{
		expression: "0 0 4 * 3",
		times: []time.Time{
			time.Date(2021, 1, 13, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 1, 20, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 1, 27, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 2, 3, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 2, 4, 0, 0, 0, 0, time.UTC),
		},
		isSet: true,
	},
}

func TestScheduleExistingJobs(t *testing.T) {
	spyHost := mocks.NewSpyHost()
	repository := mocks.NewMockRepository()
	New(repository, spyHost)
	now := time.Date(2021, 2, 4, 0, 0, 0, 0, time.UTC)
	go spyHost.TriggerTickWithTime(now)
	ScheduleExistingJobs()
	if !spyHost.WasRunJobCalled() {
		t.Fatalf("job was not called in time %v", now)
	}
}
