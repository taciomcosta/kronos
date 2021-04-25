package spynotifierservice

import (
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

type dependencyBuilder interface {
	BuildDependencies() uc.Dependencies
}

// NewSpyNotifierServiceBuilder creates a new SpyNotifierServiceBuilder
func NewSpyNotifierServiceBuilder(dependencyBuilder dependencyBuilder) *Builder {
	return &Builder{
		dependencyBuilder:  dependencyBuilder,
		spyNotifierService: newSpyNotifierService(),
	}
}

// Builder implements Builder Pattern for SpyNotifierService
type Builder struct {
	current            string
	dependencyBuilder  dependencyBuilder
	spyNotifierService *SpyNotifierService
}

// Build ...
func (s *Builder) Build() *SpyNotifierService {
	return s.spyNotifierService
}

// BuildDependencies ...
func (s *Builder) BuildDependencies() uc.Dependencies {
	return s.dependencyBuilder.BuildDependencies()
}

// Return ...
func (s *Builder) Return(vs ...interface{}) *Builder {
	s.spyNotifierService.outputs[s.current] = vs
	return s
}

// Set ...
func (s *Builder) Set(method string) *Builder {
	s.current = method
	return s
}
