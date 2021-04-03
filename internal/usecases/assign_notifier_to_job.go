package usecases

import (
	"errors"

	"github.com/taciomcosta/kronos/internal/entities"
)

// AssignNotifierToJobRequest represents request of AssignNotifierToJob
type AssignNotifierToJobRequest struct {
	JobName      string `json:"job_name"`
	NotifierName string `json:"notifier_name"`
	OnErrorOnly  bool   `json:"on_error_only"`
}

// AssignNotifierToJobResponse represents response of AssignNotifierToJob
type AssignNotifierToJobResponse struct {
	Msg string `json:"msg"`
}

// AssignNotifierToJob assigns a notifier to job
func AssignNotifierToJob(request AssignNotifierToJobRequest) (AssignNotifierToJobResponse, error) {
	response := AssignNotifierToJobResponse{}
	job, errJob := reader.FindOneJob(request.JobName)
	notifier, errNotifier := reader.FindOneNotifier(request.NotifierName)
	if errJob != nil || errNotifier != nil {
		return response, errors.New("error finding job/notifier")
	}
	assignment := entities.Assign(&job, &notifier, request.OnErrorOnly)
	err := writer.CreateAssignment(&assignment)
	if err != nil {
		return response, err
	}
	response.Msg = notifier.Name + " assigned to " + job.Name
	return response, nil
}
