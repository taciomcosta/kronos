package mocker

import uc "github.com/taciomcosta/kronos/internal/usecases"

// DependencyBuilder builds both stubs and spy dependencies
type DependencyBuilder struct {
	stubber *Stubber
}

// BuildDependencies ...
func (d *DependencyBuilder) BuildDependencies() uc.Dependencies {
	return d.stubber.BuildDependencies()
}
