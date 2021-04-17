package usecases_test

import (
	"errors"
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

func TestDescribeNotifier(t *testing.T) {
	dependencies := uc.Dependencies{
		mocks.StubSuccessWriter(),
		mocks.StubSuccessReader(),
		mocks.NewSpyHost(),
		mocks.SpyNotifierService(),
	}
	uc.New(dependencies)
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
	dependencies := uc.Dependencies{
		mocks.StubSuccessWriter(),
		mocks.StubFailingReader(),
		mocks.NewSpyHost(),
		mocks.SpyNotifierService(),
	}
	uc.New(dependencies)
	got, gotErr := uc.DescribeNotifier("list")
	want := uc.DescribeNotifierResponse{}
	wantErr := errors.New("stub-failing-reader")
	assertEqual(t, got, want)
	assertError(t, gotErr, wantErr)
}
