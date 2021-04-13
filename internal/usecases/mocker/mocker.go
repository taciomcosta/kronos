package mocker

//host := mocks.Stub().Host().RunJob().Return(executionError).Build()
//reader := mocks.Stub().Reader().FindAssignmentsByJob().Return(assignmentOnError).Build()
//writer := mocks.Stub().Writer().AlwaysSucceed().Build()
//host, wasCalled := mocks.Spy().Host().Build()

// Stub ...
func Stub() StubFn {
	return StubFn{}
}

// StubFn ...
type StubFn struct {
	//Host() stubHost
	//WriteR() stubWriter
}

// Reader ...
func (s StubFn) Reader() ReaderFn {
	outputs := make(map[string]interface{})
	returnFn := ReturnFn{outputs: outputs}
	return ReaderFn{returnFn}
}

// ReaderFn ...
type ReaderFn struct {
	returnFn ReturnFn
}

// FindAssignmentsByJob ...
func (s ReaderFn) FindAssignmentsByJob() ReturnFn {
	s.returnFn.setCurrent("FindAssignmentsByJob")
	return s.returnFn
}

// ReturnFn ...
type ReturnFn struct {
	current string
	outputs map[string]interface{}
}

// Return ...
func (r ReturnFn) Return(vs ...interface{}) ReturnFn {
	r.outputs[r.current] = vs
	return r
}

// Build ...
func (r ReturnFn) Build() StubReader {
	stub := StubReader{r}
	return stub
}

func (r ReturnFn) setCurrent(method string) {
	r.current = method
}
