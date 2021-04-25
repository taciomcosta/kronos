package mocker

import (
	"github.com/taciomcosta/kronos/internal/usecases/mocker/data"
	spyhost "github.com/taciomcosta/kronos/internal/usecases/mocker/spy_host"
	spynotifierservice "github.com/taciomcosta/kronos/internal/usecases/mocker/spy_notifier_service"
	stubreader "github.com/taciomcosta/kronos/internal/usecases/mocker/stub_reader"
	stubwriter "github.com/taciomcosta/kronos/internal/usecases/mocker/stub_writer"
)

// Dependencies ...
func Dependencies() *DependencyBuilder {
	return &DependencyBuilder{
		stubReaderBuilder:         &stubreader.Builder{},
		stubWriterBuilder:         &stubwriter.Builder{},
		spyHostBuilder:            &spyhost.Builder{},
		spyNotifierServiceBuilder: &spynotifierservice.Builder{},
	}
}

// Data ...
func Data() *data.Mocker {
	return &data.Mocker{}
}
