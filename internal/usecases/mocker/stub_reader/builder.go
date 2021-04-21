package stubreader

import (
	"github.com/taciomcosta/kronos/internal/usecases"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

type dependencyBuilder interface {
	BuildDependencies() usecases.Dependencies
}

// NewStubReaderBuilder creates a new StubReaderBuilder
func NewStubReaderBuilder(dependencyBuilder dependencyBuilder) *Builder {
	return &Builder{
		dependencyBuilder: dependencyBuilder,
		stubReader:        newStubReader(),
	}
}

// Builder implements Builder Pattern for StubReader
type Builder struct {
	current           string
	dependencyBuilder dependencyBuilder
	stubReader        *StubReader
}

// Build ...
func (s *Builder) Build() *StubReader {
	return s.stubReader
}

// BuildDependencies ...
func (s *Builder) BuildDependencies() uc.Dependencies {
	return s.dependencyBuilder.BuildDependencies()
}

// Return ...
func (s *Builder) Return(vs ...interface{}) *Builder {
	s.stubReader.outputs[s.current] = vs
	return s
}

// Set ...
func (s *Builder) Set(method string) *Builder {
	s.current = method
	return s
}
