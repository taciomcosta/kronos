package mocker

import (
	uc "github.com/taciomcosta/kronos/internal/usecases"
	"github.com/taciomcosta/kronos/internal/usecases/mocker/stub_reader"
)

// Stubber ...
type Stubber struct {
	stubReaderBuilder *stubreader.StubReaderBuilder
}

// Reader ...
func (s *Stubber) Reader() *stubreader.StubReaderBuilder {
	s.stubReaderBuilder = stubreader.NewStubReaderBuilder(s)
	return s.stubReaderBuilder
}

// BuildDependencies ...
func (s *Stubber) BuildDependencies() uc.Dependencies {
	reader := s.stubReaderBuilder.Build()
	return uc.Dependencies{
		Reader: reader,
	}
}
