package entities

// Host represents a host where jobs can be run
type Host interface {
	RunJob(job *Job) error
	GetDettachedStream() Stream
}
