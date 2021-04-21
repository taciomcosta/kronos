package mocker

import (
	"github.com/taciomcosta/kronos/internal/usecases/mocker/data"
	stubreader "github.com/taciomcosta/kronos/internal/usecases/mocker/stub_reader"
	stubwriter "github.com/taciomcosta/kronos/internal/usecases/mocker/stub_writer"
)

//import uc "github.com/taciomcosta/kronos/internal/usecases"

// TODO
//host := mocks.Stub().Host().RunJob().Return(executionError).Build()
//reader := mocks.Stub().Reader().FindAssignmentsByJob().Return(assignmentOnError).Build()
//writer := mocks.Stub().Writer().AlwaysSucceed().Build()
//host, wasCalled := mocks.Spy().Host().Build()

// Stub ...
func Stub() *Stubber {
	stubReaderBuilder := &stubreader.Builder{}
	stubWriterBuilder := &stubwriter.Builder{}
	return &Stubber{stubReaderBuilder, stubWriterBuilder}
}

// Data ...
func Data() *data.DataMocker {
	return &data.DataMocker{}
}

//func BuildDependencies() uc.Dependencies {
//return uc.Dependencies{
//Reader:          Stub().Reader().Build(),
//Writer:          Stub().Writer().Build(),
//Host:            Stub().Host().Build(),
//Notifierservice: Stub().NotifierService().Build(),
//}
//}
