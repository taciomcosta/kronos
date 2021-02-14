package sqlite

// NewWriterReader returns a Sqlite writer implementation
func NewWriterReader(name string) *WriterReader {
	newDB(name)
	return &WriterReader{}
}

// WriterReader implements usecase.Writer and usecase.Reader
type WriterReader struct{}
