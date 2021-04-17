package mocker

import uc "github.com/taciomcosta/kronos/internal/usecases"

// StubReaderBuilder ...
type StubReaderBuilder struct {
	current string
	outputs map[string]interface{}
	stubber *Stubber
}

// Build ...
func (s *StubReaderBuilder) Build() *StubReader {
	return newStubReader(s)
}

// BuildDependencies ...
func (s *StubReaderBuilder) BuildDependencies() uc.Dependencies {
	return s.stubber.BuildDependencies()
}

// Return ...
func (s *StubReaderBuilder) Return(vs ...interface{}) *StubReaderBuilder {
	s.outputs[s.current] = vs
	return s
}

// Set ...
func (s *StubReaderBuilder) Set(method string) *StubReaderBuilder {
	s.current = method
	return s
}
