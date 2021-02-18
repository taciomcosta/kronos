package usecases_test

import (
	"errors"
	"testing"

	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocks"
)

var testsDeleteJob = []struct {
	request  string
	response uc.DeleteJobResponse
	err      error
}{
	{
		request:  "name",
		response: uc.DeleteJobResponse{Msg: "name deleted"},
		err:      nil,
	},
	{
		request:  "non-existing",
		response: uc.DeleteJobResponse{},
		err:      errors.New("resource not found"),
	},
}

func TestDeleteJob(t *testing.T) {
	for _, tt := range testsDeleteJob {
		writerReader := mocks.NewStubWriterReaderNJobs(1)
		host := mocks.NewSpyHost()
		uc.New(writerReader, writerReader, host)
		got, err := uc.DeleteJob(tt.request)
		assertEqual(t, got, tt.response)
		assertError(t, err, tt.err)
	}
}
