package usecases_test

import (
	"errors"
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

func TestDescribeNotifier(t *testing.T) {
	writer := mocks.StubSuccessWriter()
	reader := mocks.StubSuccessReader()
	host := mocks.NewSpyHost()
	notifierService := mocks.NewSpyNotifierService()
	uc.New(writer, reader, host, notifierService)
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
	writer := mocks.StubSuccessWriter()
	reader := mocks.NewStubFailingReader()
	host := mocks.NewSpyHost()
	notifierService := mocks.NewSpyNotifierService()
	uc.New(writer, reader, host, notifierService)
	got, gotErr := uc.DescribeNotifier("list")
	want := uc.DescribeNotifierResponse{}
	wantErr := errors.New("stub-failing-reader")
	assertEqual(t, got, want)
	assertError(t, gotErr, wantErr)
}
