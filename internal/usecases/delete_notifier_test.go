package usecases_test

import (
	"errors"
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocker"
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
		reader:   mocker.Dependencies().Reader().Build(),
	},
	{
		request:  "non-existing",
		response: uc.DeleteNotifierResponse{},
		err:      errors.New("resource not found"),
		reader: mocker.
			Dependencies().Reader().
			Set("FindOneNotifier").
			Return(mocker.Data().Notifier().Build(), errors.New("resource not found")).
			Build(),
	},
}

func TestDeleteNotifier(t *testing.T) {
	for _, tt := range testsDeleteNotifier {
		dependencies := uc.Dependencies{
			mocker.Dependencies().Writer().Build(),
			tt.reader,
			mocker.Dependencies().Host().Build(),
			mocker.Dependencies().NotifierService().Build(),
		}
		uc.New(dependencies)
		got, err := uc.DeleteNotifier(tt.request)
		assertEqual(t, got, tt.response)
		assertError(t, err, tt.err)
	}
}
