package usecases

import "github.com/taciomcosta/kronos/internal/entities"

// CreateNotifierRequest represents the needed properties to create a notifier
type CreateNotifierRequest struct {
	Name     string                `json:"name"`
	Type     entities.NotifierType `json:"type"`
	Metadata map[string]string     `json:"metadata"`
}

// CreateNotifierResponse represents the response message of CreateNotifier
type CreateNotifierResponse struct {
	Msg string `json:"msg"`
}

// CreateNotifier creates a new notifier: slack, email, etc
func CreateNotifier(request CreateNotifierRequest) (CreateNotifierResponse, error) {
	notifier := entities.NewNotifier(
		request.Name,
		request.Type,
		request.Metadata,
	)
	err := writer.CreateNotifier(&notifier)
	if err != nil {
		return CreateNotifierResponse{}, err

	}
	return CreateNotifierResponse{Msg: notifier.Name + " created"}, nil
}
