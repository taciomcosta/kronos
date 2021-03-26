package usecases

// FindNotifiersResponse represents response of FindNotifiers usecase
type FindNotifiersResponse struct {
	Notifiers []NotifierDTO
	Count     int
}

// NotifierDTO represents a Notifier returned by FindNotifiersResponse
type NotifierDTO struct {
	Name string
	Type string
}

// FindNotifiers returns a list of all jobs.
func FindNotifiers() FindNotifiersResponse {
	return reader.FindNotifiersResponse()
}
