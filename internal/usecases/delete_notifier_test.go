package usecases_test

import (
	"errors"
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

var testsDeleteNotifier = []struct {
	request  string
	response uc.DeleteNotifierResponse
	err      error
	reader   uc.Reader
}{
	{
		request:  "myslack",
		response: uc.DeleteNotifierResponse{Msg: "myslack deleted"},
		err:      nil,
		reader:   mocks.StubSuccessReader(),
	},
	{
		request:  "non-existing",
		response: uc.DeleteNotifierResponse{},
		err:      errors.New("resource not found"),
		reader:   mocks.StubFailingReader(),
	},
}

func TestDeleteNotifier(t *testing.T) {
	for _, tt := range testsDeleteNotifier {
		writer := mocks.StubSuccessWriter()
		host := mocks.NewSpyHost()
		notifierService := mocks.SpyNotifierService()
		uc.New(writer, tt.reader, host, notifierService)
		got, err := uc.DeleteNotifier(tt.request)
		assertEqual(t, got, tt.response)
		assertError(t, err, tt.err)
	}
}
