package usecases

// UnassignNotifierFromJobRequest represents request of UnassignNotifierFromJob
type UnassignNotifierFromJobRequest struct {
	JobName      string `json:"job_name"`
	NotifierName string `json:"notifier_name"`
}

// UnassignNotifierFromJobResponse represents response of UnassignNotifierFromJob
type UnassignNotifierFromJobResponse struct {
	Msg string `json:"msg"`
}

// UnassignNotifierFromJob assigns a notifier to job
func UnassignNotifierFromJob(request UnassignNotifierFromJobRequest) (UnassignNotifierFromJobResponse, error) {
	response := UnassignNotifierFromJobResponse{}
	assignment, err := reader.FindOneAssignment(request.JobName, request.NotifierName)
	if err != nil {
		return response, err
	}
	err = writer.DeleteAssignment(&assignment)
	if err != nil {
		return response, err
	}
	response.Msg = request.NotifierName + " unassigned from " + request.JobName
	return response, nil
}
