package usecases

// DeleteJobResponse represents response of delete job usecase
type DeleteJobResponse struct {
	Msg string `json:"msg"`
}

// DeleteJob deletes a job given its name
func DeleteJob(name string) (DeleteJobResponse, error) {
	_, err := reader.FindOneJob(name)
	if err != nil {
		return DeleteJobResponse{}, err
	}
	_ = writer.DeleteJob(name)
	return DeleteJobResponse{Msg: name + " deleted."}, nil
}
