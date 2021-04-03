package entities

// Assignment represents a job assigned to a notifer
type Assignment struct {
	Job         *Job
	Notifier    *Notifier
	OnErrorOnly bool
}

// Assign assigns a job to a notifier
func Assign(job *Job, notifier *Notifier, onErrorOnly bool) Assignment {
	return Assignment{
		Job:         job,
		Notifier:    notifier,
		OnErrorOnly: onErrorOnly,
	}
}
