package usecases

// DeleteNotifierResponse represents response of delete notifier usecase
type DeleteNotifierResponse struct {
	Msg string `json:"msg"`
}

// DeleteNotifier deletes a notifier given its name
func DeleteNotifier(name string) (DeleteNotifierResponse, error) {
	_, err := reader.FindOneNotifier(name)
	if err != nil {
		return DeleteNotifierResponse{}, err
	}
	_ = writer.DeleteNotifier(name)
	return DeleteNotifierResponse{Msg: name + " deleted"}, nil
}
