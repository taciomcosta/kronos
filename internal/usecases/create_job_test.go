package usecases_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocker"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

var testsCreateJob = []struct {
	request  uc.CreateJobRequest
	response uc.CreateJobResponse
	err      error
}{
	{
		request: uc.CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "* * * * *",
		},
		response: uc.CreateJobResponse{Msg: "list created"},
		err:      nil,
	},
	{
		request: uc.CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "1 1 1 1 1",
		},
		response: uc.CreateJobResponse{Msg: "list created"},
		err:      nil,
	},
	{
		request: uc.CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "1,2,3 1,2,3 1,2,3 1,2,3 1,2,3",
		},
		response: uc.CreateJobResponse{Msg: "list created"},
		err:      nil,
	},
	{
		request: uc.CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "1,*,3 1,*,3 1,*,3 1,*,3 1,*,3",
		},
		response: uc.CreateJobResponse{Msg: "list created"},
		err:      nil,
	},
	{
		request: uc.CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "1-4 1-4 1-4 1-4 1-4",
		},
		response: uc.CreateJobResponse{Msg: "list created"},
		err:      nil,
	},
	{
		request: uc.CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "*/2 */2 */2 */2 */2",
		},
		response: uc.CreateJobResponse{Msg: "list created"},
		err:      nil,
	},
	{
		request: uc.CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "1-5/2 1-5/2 1-5/2 1-5/2 1-5/2",
		},
		response: uc.CreateJobResponse{Msg: "list created"},
		err:      nil,
	},
	{
		request: uc.CreateJobRequest{
			Name:    "list",
			Command: "ls",
			Tick:    "1-5/2,6 1-5/2,6 1-5/2,6 1-5/2,6 1-5/2,6",
		},
		response: uc.CreateJobResponse{Msg: "list created"},
		err:      nil,
	},
	{
		request:  uc.CreateJobRequest{Tick: "n * * * *"},
		response: uc.CreateJobResponse{},
		err:      errors.New("can't parse n"),
	},
	{
		request:  uc.CreateJobRequest{Tick: "* n * * *"},
		response: uc.CreateJobResponse{},
		err:      errors.New("can't parse n"),
	},
	{
		request:  uc.CreateJobRequest{Tick: "* * n * *"},
		response: uc.CreateJobResponse{},
		err:      errors.New("can't parse n"),
	},
	{
		request:  uc.CreateJobRequest{Tick: "* * * n *"},
		response: uc.CreateJobResponse{},
		err:      errors.New("can't parse n"),
	},
	{
		request:  uc.CreateJobRequest{Tick: "* * * * n"},
		response: uc.CreateJobResponse{},
		err:      errors.New("can't parse n"),
	},
	{
		request:  uc.CreateJobRequest{Tick: "60 * * * *"},
		response: uc.CreateJobResponse{},
		err:      errors.New("60 out of range 0-59"),
	},
	{
		request:  uc.CreateJobRequest{Tick: "1- * * * *"},
		response: uc.CreateJobResponse{},
		err:      errors.New("missing value in range 1-"),
	},
	{
		request:  uc.CreateJobRequest{Tick: "0-60 * * * *"},
		response: uc.CreateJobResponse{},
		err:      errors.New("0-60 out of range 0-59"),
	},
	{
		request:  uc.CreateJobRequest{Tick: "* * 0-23 * *"},
		response: uc.CreateJobResponse{},
		err:      errors.New("0-23 out of range 1-31"),
	},
	{
		request:  uc.CreateJobRequest{Tick: "* * * 1-13/2,7 *"},
		response: uc.CreateJobResponse{},
		err:      errors.New("1-13 out of range 1-12"),
	},
	{
		request:  uc.CreateJobRequest{Tick: "* * * 1-5/2,13 *"},
		response: uc.CreateJobResponse{},
		err:      errors.New("13 out of range 1-12"),
	},
	{
		request:  uc.CreateJobRequest{Tick: "10/2 * * * *"},
		response: uc.CreateJobResponse{},
		err:      errors.New("can't parse 10/2"),
	},
}

func TestCreateJob(t *testing.T) {
	dependencies := uc.Dependencies{
		mocker.Dependencies().Writer().Build(),
		mocker.Dependencies().Reader().Build(),
		mocker.Dependencies().Host().Build(),
		mocks.SpyNotifierService(),
	}
	uc.New(dependencies)
	for _, tt := range testsCreateJob {
		response, err := uc.CreateJob(tt.request)
		assertEqual(t, response, tt.response)
		assertError(t, err, tt.err)
	}
}

func assertEqual(t *testing.T, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
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

func TestCreateJobExpressionMap(t *testing.T) {
	dependencies := uc.Dependencies{
		mocker.Dependencies().Writer().Build(),
		mocker.Dependencies().Reader().Build(),
		mocker.Dependencies().Host().Build(),
		mocks.SpyNotifierService(),
	}
	uc.New(dependencies)
	for expr := range entities.SugarExpressionMap {
		request := uc.CreateJobRequest{Name: "ls", Tick: expr}
		response, err := uc.CreateJob(request)
		expectedResponse := uc.CreateJobResponse{Msg: "ls created"}
		assertEqual(t, response, expectedResponse)
		assertError(t, err, nil)
	}
}

func TestCreateJobFailingWriter(t *testing.T) {
	dependencies := uc.Dependencies{
		mocker.Dependencies().Writer().Set("CreateJob").Return(errors.New("fail")).Build(),
		mocker.Dependencies().Reader().Build(),
		mocker.Dependencies().Host().Build(),
		mocks.SpyNotifierService(),
	}
	uc.New(dependencies)
	response, err := uc.CreateJob(uc.CreateJobRequest{Tick: "* * * * *"})
	assertEqual(t, response, uc.CreateJobResponse{})
	assertError(t, err, errors.New("fail"))
}
