package sqlite

// NewWriterReader returns a Sqlite writer/reader implementation
func NewWriterReader(name string) *WriterReader {
	newDB(name)
	return &WriterReader{}
}

// NewCacheableWriterReader returns a Sqlite writer/reader implementation
// with in-memory cache
func NewCacheableWriterReader(name string) *CacheableWriterReader {
	wr := NewWriterReader(name)
	return &CacheableWriterReader{wr}
}

// WriterReader implements usecase.Writer and usecase.Reader
type WriterReader struct{}

// CacheableWriterReader implements usecase.Writer and usecase.Reader
// and adds a in-memory cache layer on top of it
type CacheableWriterReader struct {
	wr *WriterReader
}
