package mocker

import (
	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocker/stub_reader"
	"github.com/taciomcosta/kronos/internal/usecases/mocker/stub_writer"
)

// Stubber ...
type Stubber struct {
	stubReaderBuilder *stubreader.Builder
	stubWriterBuilder *stubwriter.Builder
}

// Reader ...
func (s *Stubber) Reader() *stubreader.Builder {
	s.stubReaderBuilder = stubreader.NewStubReaderBuilder(s)
	return s.stubReaderBuilder
}

// Writer ...
func (s *Stubber) Writer() *stubwriter.Builder {
	s.stubWriterBuilder = stubwriter.NewStubWriterBuilder(s)
	return s.stubWriterBuilder
}

// BuildDependencies ...
func (s *Stubber) BuildDependencies() uc.Dependencies {
	reader := s.stubReaderBuilder.Build()
	writer := s.stubWriterBuilder.Build()
	return uc.Dependencies{
		Reader: reader,
		Writer: writer,
	}
}
