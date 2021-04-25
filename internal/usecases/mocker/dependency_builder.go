package mocker

import (
	uc "github.com/taciomcosta/kronos/internal/usecases"
	spyhost "github.com/taciomcosta/kronos/internal/usecases/mocker/spy_host"
	spynotifierservice "github.com/taciomcosta/kronos/internal/usecases/mocker/spy_notifier_service"
	stubreader "github.com/taciomcosta/kronos/internal/usecases/mocker/stub_reader"
	stubwriter "github.com/taciomcosta/kronos/internal/usecases/mocker/stub_writer"
)

// DependencyBuilder builds both stubs and spy dependencies
type DependencyBuilder struct {
	stubReaderBuilder         *stubreader.Builder
	stubWriterBuilder         *stubwriter.Builder
	spyHostBuilder            *spyhost.Builder
	spyNotifierServiceBuilder *spynotifierservice.Builder
}

// BuildDependencies ...
// TODO: rename to BuildDefault()
func (d *DependencyBuilder) BuildDependencies() uc.Dependencies {
	return uc.Dependencies{
		Reader:          d.stubReaderBuilder.Build(),
		Writer:          d.stubWriterBuilder.Build(),
		Host:            d.spyHostBuilder.Build(),
		NotifierService: d.spyNotifierServiceBuilder.Build(),
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

// NotifierService ...
func (d *DependencyBuilder) NotifierService() *spynotifierservice.Builder {
	d.spyNotifierServiceBuilder = spynotifierservice.NewSpyNotifierServiceBuilder(d)
	return d.spyNotifierServiceBuilder
}
