package mocker

import uc "github.com/taciomcosta/kronos/internal/usecases"

func newStubReaderBuilder(stubber *Stubber) *StubReaderBuilder {
	return &StubReaderBuilder{
		stubber:    stubber,
		stubReader: newStubReader(),
	}
}

// StubReaderBuilder ...
type StubReaderBuilder struct {
	current    string
	stubber    *Stubber
	stubReader *StubReader
}

// Build ...
func (s *StubReaderBuilder) Build() *StubReader {
	return s.stubReader
}

// BuildDependencies ...
func (s *StubReaderBuilder) BuildDependencies() uc.Dependencies {
	return s.stubber.BuildDependencies()
}

// Return ...
func (s *StubReaderBuilder) Return(vs ...interface{}) *StubReaderBuilder {
	s.stubReader.outputs[s.current] = vs
	return s
}

// Set ...
func (s *StubReaderBuilder) Set(method string) *StubReaderBuilder {
	s.current = method
	return s
}
