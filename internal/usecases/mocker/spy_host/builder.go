package spyhost

import (
	"github.com/taciomcosta/kronos/internal/usecases"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

type dependencyBuilder interface {
	BuildDependencies() usecases.Dependencies
}

// NewSpyHostBuilder creates a new SpyHostBuilder
func NewSpyHostBuilder(dependencyBuilder dependencyBuilder) *Builder {
	return &Builder{
		dependencyBuilder: dependencyBuilder,
		spyHost:           newSpyHost(),
	}
}

// Builder implements Builder Pattern for SpyHost
type Builder struct {
	current           string
	dependencyBuilder dependencyBuilder
	spyHost           *SpyHost
}

// Build ...
func (s *Builder) Build() *SpyHost {
	return s.spyHost
}

// BuildDependencies ...
func (s *Builder) BuildDependencies() uc.Dependencies {
	return s.dependencyBuilder.BuildDependencies()
}

// Return ...
func (s *Builder) Return(vs ...interface{}) *Builder {
	s.spyHost.outputs[s.current] = vs
	return s
}

// Set ...
func (s *Builder) Set(method string) *Builder {
	s.current = method
	return s
}
