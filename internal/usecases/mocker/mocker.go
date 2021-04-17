package mocker

//import uc "github.com/taciomcosta/kronos/internal/usecases"

// TODO
//host := mocks.Stub().Host().RunJob().Return(executionError).Build()
//reader := mocks.Stub().Reader().FindAssignmentsByJob().Return(assignmentOnError).Build()
//writer := mocks.Stub().Writer().AlwaysSucceed().Build()
//host, wasCalled := mocks.Spy().Host().Build()

// Stub ...
func Stub() *Stubber {
	stubReaderBuilder := &StubReaderBuilder{}
	return &Stubber{stubReaderBuilder}
}

//func BuildDependencies() uc.Dependencies {
//return uc.Dependencies{
//Reader:          Stub().Reader().Build(),
//Writer:          Stub().Writer().Build(),
//Host:            Stub().Host().Build(),
//Notifierservice: Stub().NotifierService().Build(),
//}
//}
