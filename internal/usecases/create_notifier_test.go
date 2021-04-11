package usecases_test

import (
	"errors"
	"testing"

	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

var testsCreateNotifiers = []struct {
	request  uc.CreateNotifierRequest
	response uc.CreateNotifierResponse
	err      error
	writer   uc.Writer
}{
	{
		request: uc.CreateNotifierRequest{
			Name: "mynotifier",
			Type: entities.SlackNotifierType,
			Metadata: map[string]string{
				"auth_token":  "123",
				"channel_ids": "1,2,3",
			},
		},
		response: uc.CreateNotifierResponse{Msg: "mynotifier created"},
		err:      nil,
		writer:   mocks.StubSuccessWriter(),
	},
	{
		request: uc.CreateNotifierRequest{
			Name: "mynotifier",
			Type: entities.SlackNotifierType,
			Metadata: map[string]string{
				"auth_token":  "123",
				"channel_ids": "1,2,3",
			},
		},
		response: uc.CreateNotifierResponse{},
		err:      errors.New("StubFailingWriter"),
		writer:   mocks.NewStubFailingWriter(),
	},
	{
		request: uc.CreateNotifierRequest{
			Name:     "mynotifier",
			Type:     entities.SlackNotifierType,
			Metadata: map[string]string{},
		},
		response: uc.CreateNotifierResponse{},
		err:      errors.New("expected auth_token, channel_ids to be provided"),
		writer:   mocks.StubSuccessWriter(),
	},
}

func TestCreateNotifier(t *testing.T) {
	for _, tt := range testsCreateNotifiers {
		reader := mocks.StubSuccessReader()
		uc.New(tt.writer, reader, mocks.NewSpyHost(), mocks.NewSpyNotifierService())
		response, err := uc.CreateNotifier(tt.request)
		assertEqual(t, response, tt.response)
		assertError(t, err, tt.err)
	}
}
