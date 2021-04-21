package stubwriter

import (
	"github.com/taciomcosta/kronos/internal/usecases"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

type dependencyBuilder interface {
	BuildDependencies() usecases.Dependencies
}

// NewStubReaderBuilder creates a new StubReaderBuilder
func NewStubReaderBuilder(dependencyBuilder dependencyBuilder) *StubReaderBuilder {
	return &StubReaderBuilder{
		dependencyBuilder: dependencyBuilder,
		//stubReader:        newStubReader(),
	}
}

// StubReaderBuilder ...
type StubReaderBuilder struct {
	current           string
	dependencyBuilder dependencyBuilder
	//stubReader        *StubReader
}

// Build ...
//func (s *StubReaderBuilder) Build() *StubReader {
//return s.stubReader
//}

// BuildDependencies ...
func (s *StubReaderBuilder) BuildDependencies() uc.Dependencies {
	return s.dependencyBuilder.BuildDependencies()
}

// Return ...
//func (s *StubReaderBuilder) Return(vs ...interface{}) *StubReaderBuilder {
//s.stubReader.outputs[s.current] = vs
//return s
//}

// Set ...
func (s *StubReaderBuilder) Set(method string) *StubReaderBuilder {
	s.current = method
	return s
}
