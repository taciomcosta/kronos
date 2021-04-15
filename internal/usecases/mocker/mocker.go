package mocker

//host := mocks.Stub().Host().RunJob().Return(executionError).Build()
//reader := mocks.Stub().Reader().FindAssignmentsByJob().Return(assignmentOnError).Build()
//writer := mocks.Stub().Writer().AlwaysSucceed().Build()
//host, wasCalled := mocks.Spy().Host().Build()

// Stub ...
func Stub() *StubFn {
	readerFn := &ReaderFn{}
	return &StubFn{readerFn}
}

// StubFn ...
type StubFn struct {
	//Host() stubHost
	//WriteR() stubWriter
	readerFn *ReaderFn
}

// Reader ...
func (s *StubFn) Reader() *ReaderFn {
	outputs := make(map[string]interface{})
	s.readerFn.returnFn = &ReturnFn{outputs: outputs, readerFn: s.readerFn}
	return s.readerFn
}

// ReaderFn ...
type ReaderFn struct {
	returnFn *ReturnFn
}

// Build ...
func (s *ReaderFn) Build() *StubReader {
	return newStubReader(s.returnFn)
}

// FindAssignmentsByJob ...
func (s *ReaderFn) FindAssignmentsByJob() *ReturnFn {
	s.returnFn.setCurrent("FindAssignmentsByJob")
	return s.returnFn
}

// ReturnFn ...
type ReturnFn struct {
	current  string
	outputs  map[string]interface{}
	readerFn *ReaderFn
}

// Return ...
func (r *ReturnFn) Return(vs ...interface{}) *ReaderFn {
	r.outputs[r.current] = vs
	return r.readerFn
}

func (r *ReturnFn) setCurrent(method string) {
	r.current = method
}
