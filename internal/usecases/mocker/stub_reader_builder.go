package mocker

// StubReaderBuilder ...
type StubReaderBuilder struct {
	current string
	outputs map[string]interface{}
}

// Build ...
func (s *StubReaderBuilder) Build() *StubReader {
	return newStubReader(s)
}

// Return ...
func (s *StubReaderBuilder) Return(vs ...interface{}) *StubReaderBuilder {
	s.outputs[s.current] = vs
	return s
}

// Make ...
func (s *StubReaderBuilder) Make(method string) *StubReaderBuilder {
	s.current = method
	return s
}
