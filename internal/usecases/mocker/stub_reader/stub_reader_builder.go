package stubreader

import (
	"github.com/taciomcosta/kronos/internal/usecases"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

type DependencyBuilder interface {
	BuildDependencies() usecases.Dependencies
}

// NewStubReaderBuilder creates a new StubReaderBuilder
func NewStubReaderBuilder(dependencyBuilder DependencyBuilder) *StubReaderBuilder {
	return &StubReaderBuilder{
		dependencyBuilder: dependencyBuilder,
		stubReader:        newStubReader(),
	}
}

// StubReaderBuilder ...
type StubReaderBuilder struct {
	current           string
	dependencyBuilder DependencyBuilder
	stubReader        *StubReader
}

// Build ...
func (s *StubReaderBuilder) Build() *StubReader {
	return s.stubReader
}

// BuildDependencies ...
func (s *StubReaderBuilder) BuildDependencies() uc.Dependencies {
	return s.dependencyBuilder.BuildDependencies()
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
