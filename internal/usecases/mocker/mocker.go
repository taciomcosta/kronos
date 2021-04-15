package mocker

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

// Stubber ...
type Stubber struct {
	stubReaderBuilder *StubReaderBuilder
}

// Reader ...
func (s *Stubber) Reader() *StubReaderBuilder {
	s.stubReaderBuilder.outputs = make(map[string]interface{})
	return s.stubReaderBuilder
}
