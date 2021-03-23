package usecases_test

import (
	"errors"
	"testing"

	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

func TestCreateNotifier(t *testing.T) {
	writerReader := mocks.NewStubWriterReader()
	uc.New(writerReader, writerReader, mocks.NewSpyHost())
	request := uc.CreateNotifierRequest{
		Name: "mynotifier",
		Type: entities.SlackNotifierType,
		Metadata: map[string]string{
			"auth_token":  "123",
			"channel_ids": "1,2,3",
		},
	}
	got, err := uc.CreateNotifier(request)
	want := uc.CreateNotifierResponse{Msg: "mynotifier created"}
	assertEqual(t, got, want)
	assertError(t, err, nil)
}

func TestCreateNotifierFailure(t *testing.T) {
	writer := mocks.NewFailingWriter()
	reader := mocks.NewStubWriterReader()
	uc.New(writer, reader, mocks.NewSpyHost())
	request := uc.CreateNotifierRequest{}
	got, err := uc.CreateNotifier(request)
	want := uc.CreateNotifierResponse{}
	wantErr := errors.New("StubFailingWriter")
	assertEqual(t, got, want)
	assertError(t, err, wantErr)
}
