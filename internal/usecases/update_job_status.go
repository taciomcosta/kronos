package usecases

// UpdateJobStatusRequest represent FindExecutions request
type UpdateJobStatusRequest struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

// UpdateJobStatusResponse represents response of delete job usecase
type UpdateJobStatusResponse struct {
	Msg string `json:"msg"`
}

// UpdateJobStatus enables or disables future job executions
func UpdateJobStatus(request UpdateJobStatusRequest) (UpdateJobStatusResponse, error) {
	job, err := reader.FindOneJob(request.Name)
	if err != nil {
		return UpdateJobStatusResponse{}, err
	}
	job.Status = request.Status
	writer.UpdateJob(&job)
	if request.Status {
		return UpdateJobStatusResponse{Msg: request.Name + " enabled"}, nil
	}
	return UpdateJobStatusResponse{Msg: request.Name + " disabled"}, nil
}
