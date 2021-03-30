package usecases_test

import (
	"errors"
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

func TestDescribeNotifier(t *testing.T) {
	writer := mocks.NewStubSuccessWriter()
	reader := mocks.NewStubSuccessReader()
	host := mocks.NewSpyHost()
	uc.New(writer, reader, host)
	got, err := uc.DescribeNotifier("myslack")
	want := uc.DescribeNotifierResponse{
		Name: "myslack",
		Type: "slack",
		Metadata: map[string]string{
			"auth_token":  "123",
			"channel_ids": "1,2,3",
		},
	}
	assertEqual(t, got, want)
	assertError(t, err, nil)
}

func TestDescribeNotifierFailure(t *testing.T) {
	writer := mocks.NewStubSuccessWriter()
	reader := mocks.NewStubFailingReader()
	host := mocks.NewSpyHost()
	uc.New(writer, reader, host)
	got, gotErr := uc.DescribeNotifier("list")
	want := uc.DescribeNotifierResponse{}
	wantErr := errors.New("stub-failing-reader")
	assertEqual(t, got, want)
	assertError(t, gotErr, wantErr)
}
