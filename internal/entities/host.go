package entities

type Host interface {
	RunJob(job *Job) error
	GetDettachedStream() Stream
}
