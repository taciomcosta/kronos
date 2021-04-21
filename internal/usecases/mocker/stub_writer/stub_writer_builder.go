package stubwriter

import (
	"github.com/taciomcosta/kronos/internal/usecases"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

type dependencyBuilder interface {
	BuildDependencies() usecases.Dependencies
}

// NewStubWriterBuilder creates a new StubReaderBuilder
func NewStubWriterBuilder(dependencyBuilder dependencyBuilder) *Builder {
	return &Builder{
		dependencyBuilder: dependencyBuilder,
		stubWriter:        newStubWriter(),
	}
}

// StubWriterBuilder ...
type Builder struct {
	current           string
	dependencyBuilder dependencyBuilder
	stubWriter        *StubWriter
}

// Build ...
func (s *Builder) Build() *StubWriter {
	return s.stubWriter
}

// BuildDependencies ...
func (s *Builder) BuildDependencies() uc.Dependencies {
	return s.dependencyBuilder.BuildDependencies()
}

// Return ...
func (s *Builder) Return(vs ...interface{}) *Builder {
	s.stubWriter.outputs[s.current] = vs
	return s
}

// Set ...
func (s *Builder) Set(method string) *Builder {
	s.current = method
	return s
}
