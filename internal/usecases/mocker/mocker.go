package mocker

import (
	"github.com/taciomcosta/kronos/internal/usecases/mocker/data"
	stubreader "github.com/taciomcosta/kronos/internal/usecases/mocker/stub_reader"
	stubwriter "github.com/taciomcosta/kronos/internal/usecases/mocker/stub_writer"
)

// Stub ...
func Stub() *Stubber {
	stubReaderBuilder := &stubreader.Builder{}
	stubWriterBuilder := &stubwriter.Builder{}
	return &Stubber{stubReaderBuilder, stubWriterBuilder}
}

// Data ...
func Data() *data.DataMocker {
	return &data.DataMocker{}
}
