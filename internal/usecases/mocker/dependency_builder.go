package mocker

import (
	uc "github.com/taciomcosta/kronos/internal/usecases"
	spyhost "github.com/taciomcosta/kronos/internal/usecases/mocker/spy_host"
	stubreader "github.com/taciomcosta/kronos/internal/usecases/mocker/stub_reader"
	stubwriter "github.com/taciomcosta/kronos/internal/usecases/mocker/stub_writer"
)

// DependencyBuilder builds both stubs and spy dependencies
type DependencyBuilder struct {
	stubReaderBuilder *stubreader.Builder
	stubWriterBuilder *stubwriter.Builder
	spyHostBuilder    *spyhost.Builder
}

// BuildDependencies ...
func (d *DependencyBuilder) BuildDependencies() uc.Dependencies {
	reader := d.stubReaderBuilder.Build()
	writer := d.stubWriterBuilder.Build()
	return uc.Dependencies{
		Reader: reader,
		Writer: writer,
	}
}

// Reader ...
func (d *DependencyBuilder) Reader() *stubreader.Builder {
	d.stubReaderBuilder = stubreader.NewStubReaderBuilder(d)
	return d.stubReaderBuilder
}

// Writer ...
func (d *DependencyBuilder) Writer() *stubwriter.Builder {
	d.stubWriterBuilder = stubwriter.NewStubWriterBuilder(d)
	return d.stubWriterBuilder
}

// Host ...
func (d *DependencyBuilder) Host() *spyhost.Builder {
	d.spyHostBuilder = spyhost.NewSpyHostBuilder(d)
	return d.spyHostBuilder
}
