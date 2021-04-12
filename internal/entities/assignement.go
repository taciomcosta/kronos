package entities

// Assignment represents a job assigned to a notifer
type Assignment struct {
	Job         string
	Notifier    string
	OnErrorOnly bool
}

// Assign assigns a job to a notifier
func Assign(job *Job, notifier *Notifier, onErrorOnly bool) Assignment {
	return Assignment{
		Job:         job.Name,
		Notifier:    notifier.Name,
		OnErrorOnly: onErrorOnly,
	}
}

// ShouldNotifyExecution decides if job execution is notifiable or not
func (a Assignment) ShouldNotifyExecution(execution Execution) bool {
	return true
}
