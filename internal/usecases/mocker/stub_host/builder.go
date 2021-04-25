package stubhost

import (
	"github.com/taciomcosta/kronos/internal/usecases"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

type dependencyBuilder interface {
	BuildDependencies() usecases.Dependencies
}

// NewStubHostBuilder creates a new StubHostBuilder
func NewStubHostBuilder(dependencyBuilder dependencyBuilder) *Builder {
	return &Builder{
		dependencyBuilder: dependencyBuilder,
		stubHost:          newStubHost(),
	}
}

// Builder implements Builder Pattern for StubHost
type Builder struct {
	current           string
	dependencyBuilder dependencyBuilder
	stubHost          *StubHost
}

// Build ...
func (s *Builder) Build() *StubHost {
	return s.stubHost
}

// BuildDependencies ...
func (s *Builder) BuildDependencies() uc.Dependencies {
	return s.dependencyBuilder.BuildDependencies()
}

// Return ...
func (s *Builder) Return(vs ...interface{}) *Builder {
	s.stubHost.outputs[s.current] = vs
	return s
}

// Set ...
func (s *Builder) Set(method string) *Builder {
	s.current = method
	return s
}
