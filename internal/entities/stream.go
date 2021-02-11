package entities

import "io"

// Stream represents a stream that a job can bind to
type Stream struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

// StandardStream
// NullStream
// SpyStream
