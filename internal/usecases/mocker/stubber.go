package mocker

import uc "github.com/taciomcosta/kronos/internal/usecases"

// Stubber ...
type Stubber struct {
	stubReaderBuilder *StubReaderBuilder
}

// Reader ...
func (s *Stubber) Reader() *StubReaderBuilder {
	s.stubReaderBuilder.outputs = make(map[string]interface{})
	return s.stubReaderBuilder
}

// BuildDependencies ...
func (s *Stubber) BuildDependencies() uc.Dependencies {
	reader := s.stubReaderBuilder.Build()
	return uc.Dependencies{
		Reader: reader,
	}
}
